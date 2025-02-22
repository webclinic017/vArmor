// Copyright 2022 vArmor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package statusmanagerv1

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	v1 "k8s.io/api/core/v1"
	k8errors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"

	varmor "github.com/bytedance/vArmor/apis/varmor/v1beta1"
	varmorprofile "github.com/bytedance/vArmor/internal/profile"
	varmortypes "github.com/bytedance/vArmor/internal/types"
	varmorinterface "github.com/bytedance/vArmor/pkg/client/clientset/versioned/typed/varmor/v1beta1"
)

func getHttpBody(c *gin.Context) ([]byte, error) {
	var body []byte

	if c.Request.Body == nil {
		return body, fmt.Errorf("request body is empty")
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return body, err
	}

	return body, nil
}

// generatePolicyStatusKey build the key of StatusManager.PolicyStatuses from ProfileStatus.
//
// Its format is "VarmorClusterPolicyName" or "namespace/VarmorPolicyName".
func generatePolicyStatusKey(profileStatus *varmortypes.ProfileStatus) (string, error) {
	clusterProfileNamePrefix := fmt.Sprintf(varmorprofile.ClusterProfileNameTemplate, profileStatus.Namespace, "")
	profileNamePrefix := fmt.Sprintf(varmorprofile.ProfileNameTemplate, profileStatus.Namespace, "")
	if strings.HasPrefix(profileStatus.ProfileName, clusterProfileNamePrefix) {
		policyName := profileStatus.ProfileName[len(clusterProfileNamePrefix):]
		return policyName, nil
	} else if strings.HasPrefix(profileStatus.ProfileName, profileNamePrefix) {
		policyName := profileStatus.ProfileName[len(profileNamePrefix):]
		return profileStatus.Namespace + "/" + policyName, nil
	} else {
		return "", fmt.Errorf("profileStatus.ProfileName is illegal")
	}
}
func PolicyStatusKeyGetInfo(key string) (string, string, error) {

	keyList := strings.Split(key, "/")
	if len(keyList) == 2 {
		profileNamePrefix := fmt.Sprintf(varmorprofile.ProfileNameTemplate, keyList[0], keyList[1])
		return keyList[0], profileNamePrefix, nil
	} else if len(keyList) == 1 {
		clusterProfileNamePrefix := fmt.Sprintf(varmorprofile.ClusterProfileNameTemplate, "", keyList[0])
		return "", clusterProfileNamePrefix, nil
	} else {
		return "", key, fmt.Errorf("PolicyStatusKey is illegal")
	}

}

func generatePolicyStatusKeyWithArmorProfile(ap *varmor.ArmorProfile) (string, error) {
	clusterProfileNamePrefix := fmt.Sprintf(varmorprofile.ClusterProfileNameTemplate, ap.Namespace, "")
	profileNamePrefix := fmt.Sprintf(varmorprofile.ProfileNameTemplate, ap.Namespace, "")

	if strings.HasPrefix(ap.Name, clusterProfileNamePrefix) {
		// cluster-scope profile
		policyName := ap.Name[len(clusterProfileNamePrefix):]
		return policyName, nil
	} else if strings.HasPrefix(ap.Name, profileNamePrefix) {
		// namespace-scope profile
		policyName := ap.Name[len(profileNamePrefix):]
		return ap.Namespace + "/" + policyName, nil
	} else {
		return "", fmt.Errorf("ArmorProfile.Name is illegal")
	}
}

// generateModelingStatusKey build the key of StatusManager.ModelingStatuses from BehaviorData.
//
// Its format is "namespace/VarmorPolicyName" or "VarmorClusterPolicyName".
func generateModelingStatusKey(behaviorData *varmortypes.BehaviorData) (string, error) {
	clusterProfileNamePrefix := fmt.Sprintf(varmorprofile.ClusterProfileNameTemplate, behaviorData.Namespace, "")
	profileNamePrefix := fmt.Sprintf(varmorprofile.ProfileNameTemplate, behaviorData.Namespace, "")

	if strings.HasPrefix(behaviorData.ProfileName, clusterProfileNamePrefix) {
		// cluster-scope profile
		policyName := behaviorData.ProfileName[len(clusterProfileNamePrefix):]
		return policyName, nil
	} else if strings.HasPrefix(behaviorData.ProfileName, profileNamePrefix) {
		// namespace-scope profile
		policyName := behaviorData.ProfileName[len(profileNamePrefix):]
		return behaviorData.Namespace + "/" + policyName, nil
	} else {
		return "", fmt.Errorf("behaviorData.ProfileName is illegal")
	}
}

func newArmorProfileCondition(
	nodeName string,
	condType varmor.ArmorProfileConditionType,
	status v1.ConditionStatus,
	reason, message string) *varmor.ArmorProfileCondition {

	return &varmor.ArmorProfileCondition{
		NodeName:           nodeName,
		Type:               condType,
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

func newArmorProfileModelCondition(
	nodeName string,
	condType varmor.ArmorProfileModelConditionType,
	status v1.ConditionStatus,
	reason, message string) *varmor.ArmorProfileModelCondition {

	return &varmor.ArmorProfileModelCondition{
		NodeName:           nodeName,
		Type:               condType,
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

func UpdateVarmorPolicyStatus(
	i varmorinterface.CrdV1beta1Interface,
	vp *varmor.VarmorPolicy,
	profileName string,
	ready bool,
	phase varmor.VarmorPolicyPhase,
	condType varmor.VarmorPolicyConditionType,
	status v1.ConditionStatus,
	reason, message string) error {

	// Prepare for condition
	condition := varmor.VarmorPolicyCondition{
		Type:               condType,
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}

	regain := false
	update := func() (err error) {
		if regain {
			vp, err = i.VarmorPolicies(vp.Namespace).Get(context.Background(), vp.Name, metav1.GetOptions{})
			if err != nil {
				if k8errors.IsNotFound(err) {
					return nil
				}
				return err
			}
		}

		// Update condition
		exist := false
		switch condition.Type {
		case varmortypes.VarmorPolicyCreated:
			for i, c := range vp.Status.Conditions {
				if c.Type == varmortypes.VarmorPolicyCreated {
					condition.DeepCopyInto(&vp.Status.Conditions[i])
					exist = true
					break
				}
			}
		case varmortypes.VarmorPolicyUpdated:
			for i, c := range vp.Status.Conditions {
				if c.Type == varmortypes.VarmorPolicyUpdated {
					condition.DeepCopyInto(&vp.Status.Conditions[i])
					exist = true
					break
				}
			}
		case varmortypes.VarmorPolicyReady:
			for i, c := range vp.Status.Conditions {
				if c.Type == varmortypes.VarmorPolicyReady {
					condition.DeepCopyInto(&vp.Status.Conditions[i])
					exist = true
					break
				}
			}
		}
		if !exist {
			vp.Status.Conditions = append(vp.Status.Conditions, condition)
		}

		// Update profile name
		if profileName != "" {
			vp.Status.ProfileName = profileName
		}

		// Update status
		vp.Status.Ready = ready
		if phase != varmortypes.VarmorPolicyUnchanged {
			vp.Status.Phase = phase
		}

		_, err = i.VarmorPolicies(vp.Namespace).UpdateStatus(context.Background(), vp, metav1.UpdateOptions{})
		if err != nil {
			regain = true
		}
		return err
	}
	return retry.RetryOnConflict(retry.DefaultRetry, update)
}

func UpdateVarmorClusterPolicyStatus(
	i varmorinterface.CrdV1beta1Interface,
	vcp *varmor.VarmorClusterPolicy,
	profileName string,
	ready bool,
	phase varmor.VarmorPolicyPhase,
	condType varmor.VarmorPolicyConditionType,
	status v1.ConditionStatus,
	reason, message string) error {

	// Prepare for condition
	condition := varmor.VarmorPolicyCondition{
		Type:               condType,
		Status:             status,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}

	regain := false
	update := func() (err error) {
		if regain {
			vcp, err = i.VarmorClusterPolicies().Get(context.Background(), vcp.Name, metav1.GetOptions{})
			if err != nil {
				if k8errors.IsNotFound(err) {
					return nil
				}
				return err
			}
		}

		// Update condition
		exist := false
		switch condition.Type {
		case varmortypes.VarmorPolicyCreated:
			for i, c := range vcp.Status.Conditions {
				if c.Type == varmortypes.VarmorPolicyCreated {
					condition.DeepCopyInto(&vcp.Status.Conditions[i])
					exist = true
					break
				}
			}
		case varmortypes.VarmorPolicyUpdated:
			for i, c := range vcp.Status.Conditions {
				if c.Type == varmortypes.VarmorPolicyUpdated {
					condition.DeepCopyInto(&vcp.Status.Conditions[i])
					exist = true
					break
				}
			}
		case varmortypes.VarmorPolicyReady:
			for i, c := range vcp.Status.Conditions {
				if c.Type == varmortypes.VarmorPolicyReady {
					condition.DeepCopyInto(&vcp.Status.Conditions[i])
					exist = true
					break
				}
			}
		}
		if !exist {
			vcp.Status.Conditions = append(vcp.Status.Conditions, condition)
		}

		// Update profile name
		if profileName != "" {
			vcp.Status.ProfileName = profileName
		}

		// Update status
		vcp.Status.Ready = ready
		if phase != varmortypes.VarmorPolicyUnchanged {
			vcp.Status.Phase = phase
		}

		_, err = i.VarmorClusterPolicies().UpdateStatus(context.Background(), vcp, metav1.UpdateOptions{})
		if err != nil {
			regain = true
		}
		return err
	}
	return retry.RetryOnConflict(retry.DefaultRetry, update)
}

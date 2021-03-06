/*
 * Copyright (c) 2018 WSO2 Inc. (http:www.wso2.org) All Rights Reserved.
 *
 * WSO2 Inc. licenses this file to you under the Apache License,
 * Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http:www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"github.com/celleryio/mesh-controller/pkg/apis/mesh/v1alpha1"
	"github.com/celleryio/mesh-controller/pkg/controller"
)

func CreateTokenServiceK8sService(tokenService *v1alpha1.TokenService) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      TokenServiceK8sServiceName(tokenService),
			Namespace: tokenService.Namespace,
			Labels:    createTokenServiceLabels(tokenService),
			OwnerReferences: []metav1.OwnerReference{
				*controller.CreateTokenServiceOwnerRef(tokenService),
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				{
					Name:       tokenServiceServicePortInboundName,
					Protocol:   corev1.ProtocolTCP,
					Port:       tokenServiceServiceInboundPort,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: tokenServiceContainerInboundPort},
				},
				{
					Name:       tokenServiceServicePortOutboundName,
					Protocol:   corev1.ProtocolTCP,
					Port:       tokenServiceServiceOutboundPort,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: tokenServiceContainerOutboundPort},
				},
				{
					Name:       tokenServiceServicePortJWKSName,
					Protocol:   corev1.ProtocolTCP,
					Port:       tokenServiceContainerJWKSPort,
					TargetPort: intstr.IntOrString{Type: intstr.Int, IntVal: tokenServiceContainerJWKSPort},
				},
			},
			Selector: createTokenServiceLabels(tokenService),
		},
	}
}

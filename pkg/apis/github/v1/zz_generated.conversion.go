// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	github "github.com/nikhita/kube-custom-controller/pkg/apis/github"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_Comment_To_github_Comment,
		Convert_github_Comment_To_v1_Comment,
		Convert_v1_CommentList_To_github_CommentList,
		Convert_github_CommentList_To_v1_CommentList,
		Convert_v1_CommentSpec_To_github_CommentSpec,
		Convert_github_CommentSpec_To_v1_CommentSpec,
		Convert_v1_CommentStatus_To_github_CommentStatus,
		Convert_github_CommentStatus_To_v1_CommentStatus,
	)
}

func autoConvert_v1_Comment_To_github_Comment(in *Comment, out *github.Comment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_CommentSpec_To_github_CommentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_CommentStatus_To_github_CommentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Comment_To_github_Comment is an autogenerated conversion function.
func Convert_v1_Comment_To_github_Comment(in *Comment, out *github.Comment, s conversion.Scope) error {
	return autoConvert_v1_Comment_To_github_Comment(in, out, s)
}

func autoConvert_github_Comment_To_v1_Comment(in *github.Comment, out *Comment, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_github_CommentSpec_To_v1_CommentSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_github_CommentStatus_To_v1_CommentStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_github_Comment_To_v1_Comment is an autogenerated conversion function.
func Convert_github_Comment_To_v1_Comment(in *github.Comment, out *Comment, s conversion.Scope) error {
	return autoConvert_github_Comment_To_v1_Comment(in, out, s)
}

func autoConvert_v1_CommentList_To_github_CommentList(in *CommentList, out *github.CommentList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Items = *(*[]github.Comment)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_CommentList_To_github_CommentList is an autogenerated conversion function.
func Convert_v1_CommentList_To_github_CommentList(in *CommentList, out *github.CommentList, s conversion.Scope) error {
	return autoConvert_v1_CommentList_To_github_CommentList(in, out, s)
}

func autoConvert_github_CommentList_To_v1_CommentList(in *github.CommentList, out *CommentList, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Items = *(*[]Comment)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_github_CommentList_To_v1_CommentList is an autogenerated conversion function.
func Convert_github_CommentList_To_v1_CommentList(in *github.CommentList, out *CommentList, s conversion.Scope) error {
	return autoConvert_github_CommentList_To_v1_CommentList(in, out, s)
}

func autoConvert_v1_CommentSpec_To_github_CommentSpec(in *CommentSpec, out *github.CommentSpec, s conversion.Scope) error {
	out.Message = in.Message
	return nil
}

// Convert_v1_CommentSpec_To_github_CommentSpec is an autogenerated conversion function.
func Convert_v1_CommentSpec_To_github_CommentSpec(in *CommentSpec, out *github.CommentSpec, s conversion.Scope) error {
	return autoConvert_v1_CommentSpec_To_github_CommentSpec(in, out, s)
}

func autoConvert_github_CommentSpec_To_v1_CommentSpec(in *github.CommentSpec, out *CommentSpec, s conversion.Scope) error {
	out.Message = in.Message
	return nil
}

// Convert_github_CommentSpec_To_v1_CommentSpec is an autogenerated conversion function.
func Convert_github_CommentSpec_To_v1_CommentSpec(in *github.CommentSpec, out *CommentSpec, s conversion.Scope) error {
	return autoConvert_github_CommentSpec_To_v1_CommentSpec(in, out, s)
}

func autoConvert_v1_CommentStatus_To_github_CommentStatus(in *CommentStatus, out *github.CommentStatus, s conversion.Scope) error {
	out.Created = in.Created
	return nil
}

// Convert_v1_CommentStatus_To_github_CommentStatus is an autogenerated conversion function.
func Convert_v1_CommentStatus_To_github_CommentStatus(in *CommentStatus, out *github.CommentStatus, s conversion.Scope) error {
	return autoConvert_v1_CommentStatus_To_github_CommentStatus(in, out, s)
}

func autoConvert_github_CommentStatus_To_v1_CommentStatus(in *github.CommentStatus, out *CommentStatus, s conversion.Scope) error {
	out.Created = in.Created
	return nil
}

// Convert_github_CommentStatus_To_v1_CommentStatus is an autogenerated conversion function.
func Convert_github_CommentStatus_To_v1_CommentStatus(in *github.CommentStatus, out *CommentStatus, s conversion.Scope) error {
	return autoConvert_github_CommentStatus_To_v1_CommentStatus(in, out, s)
}

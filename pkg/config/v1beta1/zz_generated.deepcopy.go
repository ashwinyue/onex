//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GarbageCollectorControllerConfiguration) DeepCopyInto(out *GarbageCollectorControllerConfiguration) {
	*out = *in
	if in.EnableGarbageCollector != nil {
		in, out := &in.EnableGarbageCollector, &out.EnableGarbageCollector
		*out = new(bool)
		**out = **in
	}
	if in.GCIgnoredResources != nil {
		in, out := &in.GCIgnoredResources, &out.GCIgnoredResources
		*out = make([]GroupResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GarbageCollectorControllerConfiguration.
func (in *GarbageCollectorControllerConfiguration) DeepCopy() *GarbageCollectorControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(GarbageCollectorControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GenericControllerManagerConfiguration) DeepCopyInto(out *GenericControllerManagerConfiguration) {
	*out = *in
	in.LeaderElection.DeepCopyInto(&out.LeaderElection)
	out.SyncPeriod = in.SyncPeriod
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GenericControllerManagerConfiguration.
func (in *GenericControllerManagerConfiguration) DeepCopy() *GenericControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(GenericControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupResource) DeepCopyInto(out *GroupResource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupResource.
func (in *GroupResource) DeepCopy() *GroupResource {
	if in == nil {
		return nil
	}
	out := new(GroupResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MySQLConfiguration) DeepCopyInto(out *MySQLConfiguration) {
	*out = *in
	out.MaxConnectionLifeTime = in.MaxConnectionLifeTime
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MySQLConfiguration.
func (in *MySQLConfiguration) DeepCopy() *MySQLConfiguration {
	if in == nil {
		return nil
	}
	out := new(MySQLConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisConfiguration) DeepCopyInto(out *RedisConfiguration) {
	*out = *in
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisConfiguration.
func (in *RedisConfiguration) DeepCopy() *RedisConfiguration {
	if in == nil {
		return nil
	}
	out := new(RedisConfiguration)
	in.DeepCopyInto(out)
	return out
}

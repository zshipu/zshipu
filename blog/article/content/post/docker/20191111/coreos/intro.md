---
categories:
- Docker
tags:
- Docker  
keywords: 知识铺,Docker
date: 2019-11-11T22:27:21+08:00
title: Docker 从入门到实践 - CoreOS 介绍
author: 知识铺
weight: -1
---

# CoreOS 介绍

[CoreOS](https://coreos.com/) 对 Docker 甚至容器技术的发展都带来了巨大的推动作用。其提供了运行现代基础设施的特性，支持大规模服务部署，使得在基于最小化的现代操作系统上构建规模化的计算仓库成为了可能。

# CoreOS 特性

## 一个最小化操作系统

CoreOS 被设计成一个基于容器的最小化的现代操作系统。它比现有的 Linux 安装平均节省 40% 的 RAM（大约 114M ）并允许从 PXE 或 iPXE 非常快速的启动。

## 无痛更新

利用主动和被动双分区方案来更新 OS，使用分区作为一个单元而不是一个包一个包的更新。这使得每次更新变得快速，可靠，而且很容易回滚。

## Docker 容器

应用作为 Docker 容器运行在 CoreOS 上。容器以包的形式提供最大得灵活性并且可以在几毫秒启动。

## 支持集群

CoreOS 可以在一个机器上很好地运行，但是它被设计用来搭建集群。

可以通过 k8s 很容易得使应用容器部署在多台机器上并且通过服务发现把他们连接在一起。

## 分布式系统工具

内置诸如分布式锁和主选举等原生工具用来构建大规模分布式系统得构建模块。

## 服务发现

很容易定位服务在集群的那里运行并当发生变化时进行通知。它是复杂高动态集群必不可少的。
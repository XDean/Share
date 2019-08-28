# Spring Boot 入门

## 入门须知

### 你

- Java 8 熟练
- Maven 熟练
- Java Annotation 理解
- Spring 概念性了解

### 环境

- Java 8+
- Maven 3+
- Spring Boot 2.x

## 什么是Spring Boot

- Spring Boot是一个快速开发部署框架(工具集)， 帮助Java(Spring)开发者摆脱繁琐的配置
- Spring Boot不是Web Server，构建Web服务只是其诸多子项目之一
- Spring Boot核心特性
  - Spring 特性
    - 依赖注入
    - 面向切面
  - 自动配置(无XML配置文件)
  - 独立部署
  
## 现在开始

Spring Boot提供的是一个快速开发框架，其上集成了许多组件集合可以开发各类应用，其中最常见的就是Web服务。本文将主要针对Spring Boot本身和Spring Boot Web进行讲解剖析。

## Hello World

首先我们创建一个标准的maven工程，然后添加以下内容

```xml
<dependencyManagement>
    <dependencies>
        <dependency>
            <groupId>org.springframework.boot</groupId>
            <artifactId>spring-boot-starter-parent</artifactId>
            <version>2.1.7.RELEASE</version>
            <scope>import</scope>
            <type>pom</type>
        </dependency>
    </dependencies>
</dependencyManagement>

<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter</artifactId>
    </dependency>
</dependencies>
```

在`dependencyManagement`引入`spring-boot-starter-parent`来管理依赖，这样所有Spring Boot相关依赖不再需要声明版本。

*有关maven `import` scope的详细内容请参看[官方文档](https://maven.apache.org/guides/introduction/introduction-to-dependency-mechanism.html)*

引入依赖项`spring-boot-starter`，`starter`项目是一类帮助快速构建应用的工具项目，其包括了构建该类应用的必备和首选依赖，例如`spring-boot-starter-web`就包括了spring-boot, web, json, validator, tomcat等依赖。


# E-Learning Management System based on Go #

## Contents

1. [Description](#description)
1. [Objectives](#objectives)
1. [Motivation](#motivation)
1. [Tech Stack](#tech-stack)
1. [Deployment](#deployment)
1. [Product](#product)
1. [Credits](#credits)

## Description
[⬆ Back to the conents](#contents)

Backend (capstone?) project for a distributed E-Learning Management System (LMS) developed fully in Go-Lang.

## Objectives
[⬆ Back to the conents](#contents)

- Creating an easy to maintain backend project.
- Using go-lang and learning it from scratch with this project.
- Create a Learning Management System that can actually be deployed.
- Creating a Postman collection for every user to use.

## Motivation
[⬆ Back to the conents](#contents)

I wanted to learn Go for so long, like, years even, but it was always I said I'd do later on, this is later on and I'm doing it. Part of the idea was, not only seeing Go on my feed, but this particular video, [Rust or Go for my next project? WHAT TO CHOOSE? (as a senior intern engineer)](https://www.youtube.com/watch?v=LbmvbXPj8Fs).

Go is a programming language developed and maintained by Google, a quite popular bussiness or so they say, it was one of my (technical) objectives for 2022. I've head good things about it's concurrency and it's simplicity, how it manages packages/modules and dependencies, how easy-to-understand the codebase can be and it's efficiency. I also really like the idea that it is compiled, it's been a long time since I have worked with compiled languages, and, as far as it goes, Java is not on my priorities, for the moment being that is.

Another plus is that it's strongly typed. And, finally, obviously, it offers it's own testing packages.

## Tech Stack
[⬆ Back to the conents](#contents)

- **Golang**, as aforementioned, a popuplar compiled language developed by Google that allows concurrency and offers a type system.
- **MySQL**, a relational database for SQL, and it's the one I'm the most comfortable with.

## Deployment

I'll be using Github Actions as my main source of CI/CD configuration, the idea is for it to be _ready_ to be deployed to a Cloud Computing Hyperscaler, not necessarily that it is going to be deployed, even though I'd like to test it on the cloud.

I'll also be using Docker + Kubernetes for the container and orchestration. This part may suffer slight changes along the way.

## Product
[⬆ Back to the conents](#contents)

I want to make an e-learning that's easy to maintain, easily scalable and _"blazingly fast"_.

My product should allow for an administrator allowing different coordinators to create different courses composed of several modules imparted by one teacher to many students so that they (the students) can complete the proposed quizzes, activities and complete the required/optional activities that may be deemed necessary so that it's as successful as possible.

With all the information generated via those actions, coming from (small) data science projects that I've been doing lately, dashboards may be composed so that reports and queries can help fullfill the goals the coordinators may have for their courses and teachers, coordinators and administrators can actually improve their content.

Since it's going to be done on my free time with no other clear idea than playing around and learn this new technology (Golang), there's no defined Product Roadmap no ticket management, no management at all to be clear. It may be implemented later down the line.

## Credits
[⬆ Back to the conents](#contents)

Using Go-Lang motivated by this YouTube video: [Rust or Go for my next project? WHAT TO CHOOSE? (as a senior intern engineer)](https://www.youtube.com/watch?v=LbmvbXPj8Fs).

Project idea from [20 Exciting Software Development Project Ideas & Topics for Beginners [2022]](https://www.upgrad.com/blog/software-development-project-ideas-topics-for-beginners/#18_e-Learning_platform).
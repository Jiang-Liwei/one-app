package factories

import (
    "forum/app/models/project"
    "forum/pkg/helpers"

    "github.com/bxcodec/faker/v3"
)

func MakeProjects(count int) []project.Project {

    var objs []project.Project

    // 设置唯一性，如 Project 模型的某个字段需要唯一，即可取消注释
    // faker.SetGenerateUniqueValues(true)

    for i := 0; i < count; i++ {
        projectModel := project.Project{
            FIXME()
        }
        objs = append(objs, projectModel)
    }

    return objs
}
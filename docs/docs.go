// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/failed": {
            "get": {
                "description": "失败测试接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试相关接口"
                ],
                "summary": "Failed测试接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/success": {
            "get": {
                "description": "成功测试接口",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "测试相关接口"
                ],
                "summary": "Success测试接口",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysMenu/add": {
            "post": {
                "description": "新增菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "新增菜单",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddSysMenuDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysMenu/delete": {
            "delete": {
                "description": "根据id删除菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "根据id删除菜单",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysMenuIdDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysMenu/info": {
            "get": {
                "description": "根据id查询菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "根据id查询菜单",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "菜单id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysMenu/list": {
            "get": {
                "description": "查询菜单列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "查询菜单列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "菜单名称",
                        "name": "menuName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "菜单状态",
                        "name": "menuStatus",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysMenu/update": {
            "put": {
                "description": "修改菜单",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "菜单相关接口"
                ],
                "summary": "修改菜单",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysMenuDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysRole/add": {
            "post": {
                "description": "新增角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "新增角色",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateSysRoleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysRole/delete": {
            "delete": {
                "description": "根据id删除角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "根据id删除角色",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysRoleIdDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysRole/info": {
            "get": {
                "description": "根据id查询角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "根据id查角色",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "角色id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysRole/list": {
            "get": {
                "description": "分页查询角色列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "分页查询角色列表",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "分页数量",
                        "name": "pageNum",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页大小",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色名称",
                        "name": "roleName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "角色状态",
                        "name": "roleStatus",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "开始时间",
                        "name": "beginTime",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "结束时间",
                        "name": "endTime",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysRole/update": {
            "put": {
                "description": "修改角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "修改角色",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysRoleDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysRole/updateStatus": {
            "put": {
                "description": "修改角色状态",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "修改角色状态",
                "parameters": [
                    {
                        "description": "data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateSysRoleStatusDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        },
        "/api/sysRole/vo/list": {
            "get": {
                "description": "查询角色下拉列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色相关接口"
                ],
                "summary": "查询角色下拉列表",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/result.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddSysMenuDto": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "menuName": {
                    "type": "string"
                },
                "menuStatus": {
                    "type": "integer"
                },
                "menuType": {
                    "type": "integer"
                },
                "parentId": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.CreateSysRoleDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "roleKey": {
                    "type": "string"
                },
                "roleName": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.SysMenuIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.SysRoleIdDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysMenuDto": {
            "type": "object",
            "properties": {
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "menuName": {
                    "type": "string"
                },
                "menuStatus": {
                    "type": "integer"
                },
                "menuType": {
                    "type": "integer"
                },
                "parentId": {
                    "type": "integer"
                },
                "sort": {
                    "type": "integer"
                },
                "url": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.UpdateSysRoleDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "roleKey": {
                    "type": "string"
                },
                "roleName": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.UpdateSysRoleStatusDto": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "result.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "后台系统",
	Description:      "后台API接口文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

<template>
    <div>
        <el-row class="top">
            <el-row class="left">
                <el-button size="small" @click="add" type="primary">添加</el-button>
            </el-row>
            <el-form :inline="true" class="demo-form-inline right">
                <el-form-item>
                    <el-input placeholder="课程名称" size="small"></el-input>
                </el-form-item>
                <el-form-item>
                    <el-button size="small" type="primary">查询</el-button>
                </el-form-item>
            </el-form>
        </el-row>
        <el-table
            :data="tableData"
            style="width: 100%">
            <el-table-column
                fixed
                align="center"
                prop="name"
                label="课程名称"
                width="200"
            >
            </el-table-column>
            <el-table-column
                align="center"
                prop="price"
                label="价格"
                width="200"
                >
                <template slot-scope="data">
                    <span class="red">{{ data.row.price }}</span>
                </template>
            </el-table-column>
            <el-table-column
                align="center"
                prop="description"
                label="简介"
                width="300"
            >
            </el-table-column>
            <el-table-column
                align="center"
                prop="is_recommend"
                label="是否推荐"
                width="200"
            >
                <template slot-scope="data">
                    <span style="margin-left: 10px">{{ data.row.is_recommend ? "是" : "否" }}</span>
                </template>
            </el-table-column>
            <el-table-column
                align="center"
                prop="is_free"
                label="是否免费"
                width="300"
            >
                <template slot-scope="data">
                    <span style="margin-left: 10px">{{ data.row.is_free ? "是" : "否" }}</span>
                </template>
            </el-table-column>
            <el-table-column
                align="center"
                prop="created_at"
                label="创建时间"
                width="200"
            >
            </el-table-column>
            <el-table-column
                align="center"
                prop="updated_at"
                label="修改时间"
                width="200"
            >
            </el-table-column>
            <el-table-column
                align="center"
                fixed="right"
                label="操作"
                width="150">
                    <template slot-scope="scope">
                        <el-button
                            @click.native.prevent="editRow(scope.row.id)"
                            type="text"
                            size="small">
                            修改
                        </el-button>
                        <el-button
                            @click.native.prevent="section(scope.row.id)"
                            type="text"
                            size="small">
                            章节
                        </el-button>
                        <el-button
                            @click.native.prevent="deleteRow(scope.row.id)"
                            type="text"
                            size="small">
                            移除
                        </el-button>
                    </template>
                </el-table-column>
        </el-table>
        <el-pagination
            background
            layout="prev, pager, next"
            :total="1000">
        </el-pagination>
    </div>
</template>

<script>
import cource from "../../api/course";
import c from "../../lib/const"

export default {
    data() {
        return {
            tableData: []
        }
    },
    components: {
    },
    mounted() {
        this.list()
    },
    methods: {
        deleteRow(id) {
            this.$confirm("确定要删除吗？", "提示", {type: "warning"}).then(() => {
                cource.delete(id).then((res) => {
                    if (res.code == c.success) {
                        this.$message({
                            message: "删除成功",
                            type: "success"
                        })
                        this.list()
                    }
                }).catch(() => {
                    this.$message.error("删除失败")
                })
            })
        },
        editRow(id) {
            this.$router.push({name:"课程修改", params: {id}})
        },
        add() {
            this.$router.push({name:"课程添加"})
        },
        list() {
           cource.list(
               {},
               (res) => {
                   this.tableData = res.data
               },
               (res) => {
                   window.console.log(res.data)
               }
            )
        },
        section(cource_id) {
            this.$router.push({name:"章节列表", params:{cid: cource_id}})
        }
    }
}
</script>

<style scoped lang="scss">
    @import "@/assets/css/variables.sass";
    .el-table td, .el-table th{
        padding: 0px!important;
    }
    .el-pagination {
        margin: 10px 0 0;
        text-align: right;
    }
    .top{
        .left {
            float: left;
        }
        .right {
            float: right;
            .el-form-item {
                margin-right: 0px;
                margin-left: 10px;
            }
        }
    }
    .red {
        color: $red;
    }
</style>
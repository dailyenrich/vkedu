<template>
    <div>
        <el-form label-position="left" label-width="80px" :model="section">
            <el-form-item label="章节名称">
                <el-input class="input-item" v-model="section.name"></el-input>
            </el-form-item>
            <el-form-item label="章节描述">
                <el-input type="textarea" v-model="section.description" class="desc"></el-input>
            </el-form-item>
        </el-form>
        <el-row>
            <el-button type="primary" @click="update" style="margin-left:80px">保存</el-button>
            <el-button type="info" @click="reset">重置</el-button>
            <el-button type="info" @click="back">返回</el-button>
        </el-row>
    </div>
</template>

<script>
import section from "../../api/section";
import c from "../../lib/const";

export default {
    data() {
        return {
            section: {
                course_id: null
            },
            old: {}
        }
    },
    mounted() {
        this.getOne()
    },
    methods: {
        getOne: function() {
            let _this = this
            section.getOneByID(this.$route.params['id']).then(function(res) {
                if (res.code == c.success) {
                    _this.section = res.data
                    _this.old = res.data
                }
            }).catch(function() {
                _this.$message.error("获取数据失败")
            })
        },
        update() {
            section.edit(this.section).then((res) => {
                if (res.code == c.success) {
                    this.$message({
                        message: "保存成功",
                        type: "success"
                    })
                    this.back()
                } else {
                    this.$message.error(res.message)
                }
            }).catch(() => {
                this.$message.error("保存失败")
            })
        },
        reset: function() {
            this.section = this.old
        },
        back: function() {
            this.$router.go(-1)
        }
    }
}
</script>

<style lang="scss" scoped>
    @import "@/assets/css/variables.sass"; 
    .input-item {
        width: 500px;
    }    
    .desc textarea {
        height: 200px!important;
    }
</style>
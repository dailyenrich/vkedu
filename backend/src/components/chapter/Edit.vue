<template>
    <div>
        <el-form label-position="left" label-width="80px" :model="chapter">
            <el-form-item label="章节名称">
                <el-input class="input-item" v-model="chapter.name"></el-input>
            </el-form-item>
            <el-form-item label="章节描述">
                <el-input type="textarea" v-model="chapter.description" class="desc"></el-input>
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
import chapter from "../../api/chapter";
import c from "../../lib/const";

export default {
    data() {
        return {
            chapter: {
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
            chapter.getOneByID(this.$route.params['id']).then(function(res) {
                if (res.code == c.success) {
                    _this.chapter = res.data
                    _this.old = res.data
                }
            }).catch(function() {
                _this.$message.error("获取数据失败")
            })
        },
        update() {
            chapter.edit(this.chapter).then((res) => {
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
            this.chapter = this.old
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
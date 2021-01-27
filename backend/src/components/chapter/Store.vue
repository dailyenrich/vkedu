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
            <el-button type="primary" @click="store" style="margin-left:80px">添加</el-button>
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
        }
    },
    mounted() {
        this.chapter.course_id = parseInt(this.$route.params['cid'])
    },
    methods: {
        store: function() {
            chapter.store(this.chapter).then((res) => {
                    if (res.code == c.success) {
                        this.$message({
                            message: "添加成功",
                            type: "success"
                        })
                        this.reset()
                    } else {
                        this.$message.error("添加失败")
                    }
                }).catch(() => {
                    this.$message.error("添加失败")
                })
        },
        reset: function() {
            this.chapter = {
                course_id: parseInt(this.$route.params['cid'])
            }
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
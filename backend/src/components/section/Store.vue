<template>
    <div>
        <el-form label-position="left" label-width="80px" :model="section">
            <el-form-item label="小节名称">
                <el-input class="input-item" v-model="section.name"></el-input>
            </el-form-item>
            <el-form-item label="是否免费">
                <el-radio v-model="section.is_free" :label="false">收费</el-radio>
                <el-radio v-model="section.is_free" :label="true">免费</el-radio>
            </el-form-item>
            <el-form-item label="小节描述">
                <el-input type="textarea" v-model="section.description" class="desc"></el-input>
            </el-form-item>
            <el-form-item label="小节详情">
                <MarkdownEditor mode="markdown" v-model="section.detail" />
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
import section from "../../api/section";
import c from "../../lib/const";
import MarkdownEditor from "../markdown/index";

export default {
    data() {
        return {
            section: {
                chapter_id: null
            },
        }
    },
    components: {
        MarkdownEditor
    },
    mounted() {
        this.section.chapter_id = parseInt(this.$route.params['cid'])
    },
    methods: {
        store: function() {
            section.store(this.section).then((res) => {
                window.console.log(res)
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
            this.section = {
                chapter_id: parseInt(this.$route.params['cid'])
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
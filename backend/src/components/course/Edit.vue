<template>
    <div>
        <el-form label-position="left" label-width="80px" :model="cource">
            <el-form-item label="名称">
                <el-input class="input-item" v-model="cource.name"></el-input>
            </el-form-item>
            <el-form-item label="是否免费">
                <el-radio v-model="cource.is_free" :label="false">收费</el-radio>
                <el-radio v-model="cource.is_free" :label="true">免费</el-radio>
            </el-form-item>
            <el-form-item label="是否热门">
                <el-radio v-model="cource.is_recommend" :label="false">不推荐</el-radio>
                <el-radio v-model="cource.is_recommend" :label="true">推荐</el-radio>
            </el-form-item>
            <el-form-item label="价格">
                <el-input-number v-model="cource.price" :precision="2" :step="0.01"></el-input-number>
            </el-form-item>
            <el-form-item label="课程描述">
                <el-input type="textarea" v-model="cource.description" class="desc"></el-input>
            </el-form-item>
            <el-form-item label="课程详情">
                <MarkdownEditor mode="markdown" v-model="cource.detail" />
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
import MarkdownEditor from "../markdown/index";
import cource from "../../api/course";
import c from "../../lib/const"

export default {
    data() {
        return {
            cource: {
                is_free: false,
                is_recommend: false,
                price: 0
            },
            old: {},
            content: "",
            num: 0
        }
    },
    mounted() {
        this.getCource()
    },
    components: {
        MarkdownEditor
    },
    methods: {
        getCource: function() {
            let _this = this
            cource.getCourceByID(this.$route.params['id']).then(function(res) {
                if (res.code == c.success) {
                    _this.cource = res.data
                    _this.old = res.data
                }
            }).catch(function() {
                _this.$message.error("获取数据失败")
            })
        },
        update() {
            cource.edit(this.cource).then((res) => {
                if (res.code == c.success) {
                    this.$message({
                        message: "保存成功",
                        type: "success"
                    })
                    this.back()
                }
            }).catch(() => {
                this.$message.error("保存失败")
            })
        },
        reset: function() {
            this.cource = this.old
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
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
            <el-button type="primary" @click="store" style="margin-left:80px">添加</el-button>
            <el-button type="info" @click="reset">重置</el-button>
            <el-button type="info" @click="back">返回</el-button>
        </el-row>
    </div>
</template>

<script>
import MarkdownEditor from "../markdown/index";
import courese from "../../api/course";

export default {
    data() {
        return {
            cource: {
                is_free: false,
                is_recommend: false,
                price: 0
            },
            content: "# ok\n## ol",
            num: 0
        }
    },
    components: {
        MarkdownEditor
    },
    methods: {
        store: function() {
            let _this = this
            courese.store(this.cource, function() {
                _this.$message({
                    message: '课程添加成功',
                    type: 'success'
                });
                _this.reset()
            }, function() {
                _this.$message.error('课程添加失败');
            })
        },
        reset: function() {
            this.cource = {}
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

<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="投票项目名称:" prop="voteName">
          <el-input v-model="formData.voteName" :clearable="true"  placeholder="请输入投票项目名称" />
       </el-form-item>
        <el-form-item label="投票项目图片:" prop="votePicture">
          <SelectImage v-model="formData.votePicture" file-type="image"/>
       </el-form-item>
        <el-form-item label="投票是否通过审核:" prop="voteIsPass">
          <el-switch v-model="formData.voteIsPass" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="投票项目是否展示:" prop="voteIsShow">
          <el-switch v-model="formData.voteIsShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="公链id:" prop="publicChainId">
           <el-select v-model="formData.publicChainId" placeholder="请选择公链id" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="投票总数:" prop="voteNum">
          <el-input v-model.number="formData.voteNum" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="投票项目链接:" prop="voteUrl">
          <el-input v-model="formData.voteUrl" :clearable="true"  placeholder="请输入投票项目链接" />
       </el-form-item>
        <el-form-item label="简介:" prop="brief">
          <el-input v-model="formData.brief" :clearable="true"  placeholder="请输入简介" />
       </el-form-item>
        <el-form-item label="详情描述:" prop="detail">
          <RichEdit v-model="formData.detail"/>
       </el-form-item>
        <el-form-item label="x链接:" prop="xLink">
          <el-input v-model="formData.xLink" :clearable="true"  placeholder="请输入x链接" />
       </el-form-item>
        <el-form-item label="tg链接:" prop="tgLink">
          <el-input v-model="formData.tgLink" :clearable="true"  placeholder="请输入tg链接" />
       </el-form-item>
        <el-form-item label="discord链接:" prop="discordLink">
          <el-input v-model="formData.discordLink" :clearable="true"  placeholder="请输入discord链接" />
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createVote,
  updateVote,
  findVote
} from '@/api/Vote/vote'

defineOptions({
    name: 'VoteForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const PublicChainOptions = ref([])
const formData = ref({
            voteName: '',
            votePicture: "",
            voteIsPass: false,
            voteIsShow: false,
            publicChainId: '',
            voteNum: undefined,
            voteUrl: '',
            brief: '',
            detail: '',
            xLink: '',
            tgLink: '',
            discordLink: '',
        })
// 验证规则
const rule = reactive({
               voteName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               votePicture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               publicChainId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findVote({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    PublicChainOptions.value = await getDictFunc('PublicChain')
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createVote(formData.value)
               break
             case 'update':
               res = await updateVote(formData.value)
               break
             default:
               res = await createVote(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>


<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="标题:" prop="title">
          <el-input v-model="formData.title" :clearable="true"  placeholder="请输入标题" />
       </el-form-item>
        <el-form-item label="作者:" prop="author">
          <el-input v-model="formData.author" :clearable="true"  placeholder="请输入作者" />
       </el-form-item>
        <el-form-item label="正文:" prop="content">
          <RichEdit v-model="formData.content"/>
       </el-form-item>
        <el-form-item label="是否展示:" prop="is_Show">
          <el-switch v-model="formData.is_Show" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="资讯分类:" prop="publicChainId">
           <el-select v-model="formData.publicChainId" placeholder="请选择资讯分类" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in 咨询二级标题Options" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="封面图:" prop="picture">
          <SelectImage v-model="formData.picture" file-type="image"/>
       </el-form-item>
        <el-form-item label="作者头像:" prop="authorHeadImg">
          <SelectImage v-model="formData.authorHeadImg" file-type="image"/>
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
  createInformation,
  updateInformation,
  findInformation
} from '@/api/Information/information'

defineOptions({
    name: 'InformationForm'
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
const 咨询二级标题Options = ref([])
const formData = ref({
            title: '',
            author: '',
            content: '',
            is_Show: false,
            publicChainId: '',
            picture: "",
            authorHeadImg: "",
        })
// 验证规则
const rule = reactive({
               publicChainId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               picture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               authorHeadImg : [{
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
      const res = await findInformation({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    咨询二级标题Options.value = await getDictFunc('咨询二级标题')
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
               res = await createInformation(formData.value)
               break
             case 'update':
               res = await updateInformation(formData.value)
               break
             default:
               res = await createInformation(formData.value)
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

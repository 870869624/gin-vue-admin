
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户名:" prop="username">
          <el-input v-model="formData.username" :clearable="true"  placeholder="请输入用户名" />
       </el-form-item>
        <el-form-item label="密码:" prop="password">
          <el-input v-model="formData.password" :clearable="true"  placeholder="请输入密码" />
       </el-form-item>
        <el-form-item label="邮件:" prop="email">
          <el-input v-model="formData.email" :clearable="true"  placeholder="请输入邮件" />
       </el-form-item>
        <el-form-item label="电话号码:" prop="phoneNumber">
          <el-input v-model="formData.phoneNumber" :clearable="true"  placeholder="请输入电话号码" />
       </el-form-item>
        <el-form-item label="头像:" prop="avatar">
          <SelectImage v-model="formData.avatar" file-type="image"/>
       </el-form-item>
        <el-form-item label="个人简介:" prop="brief">
          <el-input v-model="formData.brief" :clearable="true"  placeholder="请输入个人简介" />
       </el-form-item>
        <el-form-item label="年龄:" prop="age">
          <el-input v-model.number="formData.age" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="性别:" prop="gender">
           <el-select v-model="formData.gender" placeholder="请选择性别" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="住址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入住址" />
       </el-form-item>
        <el-form-item label="小狐狸钱包地址:" prop="metaMask">
          <el-input v-model="formData.metaMask" :clearable="true"  placeholder="请输入小狐狸钱包地址" />
       </el-form-item>
        <el-form-item label="TP钱包地址:" prop="tokenPocket">
          <el-input v-model="formData.tokenPocket" :clearable="true"  placeholder="请输入TP钱包地址" />
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
  createUsers,
  updateUsers,
  findUsers
} from '@/api/Users/users'

defineOptions({
    name: 'UsersForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const genderOptions = ref([])
const formData = ref({
            username: '',
            password: '',
            email: '',
            phoneNumber: '',
            avatar: "",
            brief: '',
            age: undefined,
            gender: '',
            address: '',
            metaMask: '',
            tokenPocket: '',
        })
// 验证规则
const rule = reactive({
               username : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               password : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               email : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               phoneNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gender : [{
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
      const res = await findUsers({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
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
               res = await createUsers(formData.value)
               break
             case 'update':
               res = await updateUsers(formData.value)
               break
             default:
               res = await createUsers(formData.value)
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


<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAt">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
      </el-form-item>
      
        <el-form-item label="空投名称" prop="airdropName">
         <el-input v-model="searchInfo.airdropName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="空投项目分发日期" prop="distributionDate">
         <el-input v-model="searchInfo.distributionDate" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="空投项目价值" prop="airdropValue">
            
             <el-input v-model.number="searchInfo.airdropValue" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="空投项目结束日期" prop="airdropEndtime">
            
            <template #label>
            <span>
              空投项目结束日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startAirdropEndtime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endAirdropEndtime ? time.getTime() > searchInfo.endAirdropEndtime.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endAirdropEndtime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startAirdropEndtime ? time.getTime() < searchInfo.startAirdropEndtime.getTime() : false"></el-date-picker>
        </el-form-item>
        
           <el-form-item label="公链id" prop="publicChainId">
            <el-select v-model="searchInfo.publicChainId" clearable placeholder="请选择" @clear="()=>{searchInfo.publicChainId=undefined}">
              <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
            </el-form-item>
            <el-form-item label="是否通过审核" prop="airdropIsPass">
            <el-select v-model="searchInfo.airdropIsPass" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
            <el-form-item label="是否展示" prop="airdropIsShow">
            <el-select v-model="searchInfo.airdropIsShow" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            <ExportTemplate  template-id="Airdrop_Airdrop" />
            <ExportExcel  template-id="Airdrop_Airdrop" />
            <ImportExcel  template-id="Airdrop_Airdrop" @on-success="getTableData" />
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="日期" prop="createdAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
          <el-table-column align="left" label="空投名称" prop="airdropName" width="120" />
          <el-table-column label="空投项目图片" prop="airdropPicture" width="200">
              <template #default="scope">
                <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.airdropPicture)" fit="cover"/>
              </template>
          </el-table-column>
          <el-table-column align="left" label="空投项目价值" prop="airdropValue" width="120" />
         <el-table-column align="left" label="空投项目结束日期" prop="airdropEndtime" width="180">
            <template #default="scope">{{ formatDate(scope.row.airdropEndtime) }}</template>
         </el-table-column>
         <el-table-column align="left" label="空投项目分发日期" prop="distributionDate" width="180">
            <template #default="scope">{{ formatDate(scope.row.distributionDate) }}</template>
         </el-table-column>
        <el-table-column align="left" label="公链id" prop="publicChainId" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.publicChainId,PublicChainOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="是否通过审核" prop="airdropIsPass" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.airdropIsPass) }}</template>
        </el-table-column>
        <el-table-column align="left" label="是否展示" prop="airdropIsShow" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.airdropIsShow) }}</template>
        </el-table-column>
          <el-table-column align="left" label="空有项目网址" prop="airdropUrl" width="120" />
          <el-table-column align="left" label="简介" prop="brief" width="120" />
                      <el-table-column label="详情描述" prop="detail" width="200">
                         <template #default="scope">
                            [富文本内容]
                         </template>
                      </el-table-column>
          <el-table-column align="left" label="x链接" prop="xLink" width="120" />
          <el-table-column align="left" label="TG链接" prop="tgLink" width="120" />
          <el-table-column align="left" label="discord链接" prop="discordLink" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateAirdropFunc(scope.row)">编辑</el-button>
            <el-button  type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="空投名称:"  prop="airdropName" >
              <el-input v-model="formData.airdropName" :clearable="true"  placeholder="请输入空投名称" />
            </el-form-item>
            <el-form-item label="空投项目图片:"  prop="airdropPicture" >
                <SelectImage
                 v-model="formData.airdropPicture"
                 file-type="image"
                />
            </el-form-item>
            <el-form-item label="空投项目价值:"  prop="airdropValue" >
              <el-input-number v-model="formData.airdropValue"  style="width:100%" :precision="2" :clearable="true"  />
            </el-form-item>
            <el-form-item label="空投项目结束日期:"  prop="airdropEndtime" >
              <el-date-picker v-model="formData.airdropEndtime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="空投项目分发日期:"  prop="distributionDate" >
              <el-date-picker v-model="formData.distributionDate" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="公链id:"  prop="publicChainId" >
              <el-select v-model="formData.publicChainId" placeholder="请选择公链id" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="是否通过审核:"  prop="airdropIsPass" >
              <el-switch v-model="formData.airdropIsPass" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="是否展示:"  prop="airdropIsShow" >
              <el-switch v-model="formData.airdropIsShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="空有项目网址:"  prop="airdropUrl" >
              <el-input v-model="formData.airdropUrl" :clearable="true"  placeholder="请输入空有项目网址" />
            </el-form-item>
            <el-form-item label="简介:"  prop="brief" >
              <el-input v-model="formData.brief" :clearable="true"  placeholder="请输入简介" />
            </el-form-item>
            <el-form-item label="详情描述:"  prop="detail" >
              <RichEdit v-model="formData.detail"/>
            </el-form-item>
            <el-form-item label="x链接:"  prop="xLink" >
              <el-input v-model="formData.xLink" :clearable="true"  placeholder="请输入x链接" />
            </el-form-item>
            <el-form-item label="TG链接:"  prop="tgLink" >
              <el-input v-model="formData.tgLink" :clearable="true"  placeholder="请输入TG链接" />
            </el-form-item>
            <el-form-item label="discord链接:"  prop="discordLink" >
              <el-input v-model="formData.discordLink" :clearable="true"  placeholder="请输入discord链接" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="空投名称">
                        {{ detailFrom.airdropName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="空投项目图片">
                            <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailFrom.airdropPicture)" :src="getUrl(detailFrom.airdropPicture)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="空投项目价值">
                        {{ detailFrom.airdropValue }}
                    </el-descriptions-item>
                    <el-descriptions-item label="空投项目结束日期">
                        {{ detailFrom.airdropEndtime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="空投项目分发日期">
                        {{ detailFrom.distributionDate }}
                    </el-descriptions-item>
                    <el-descriptions-item label="公链id">
                        {{ detailFrom.publicChainId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="是否通过审核">
                        {{ detailFrom.airdropIsPass }}
                    </el-descriptions-item>
                    <el-descriptions-item label="是否展示">
                        {{ detailFrom.airdropIsShow }}
                    </el-descriptions-item>
                    <el-descriptions-item label="空有项目网址">
                        {{ detailFrom.airdropUrl }}
                    </el-descriptions-item>
                    <el-descriptions-item label="简介">
                        {{ detailFrom.brief }}
                    </el-descriptions-item>
                    <el-descriptions-item label="详情描述">
                        {{ detailFrom.detail }}
                    </el-descriptions-item>
                    <el-descriptions-item label="x链接">
                        {{ detailFrom.xLink }}
                    </el-descriptions-item>
                    <el-descriptions-item label="TG链接">
                        {{ detailFrom.tgLink }}
                    </el-descriptions-item>
                    <el-descriptions-item label="discord链接">
                        {{ detailFrom.discordLink }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createAirdrop,
  deleteAirdrop,
  deleteAirdropByIds,
  updateAirdrop,
  findAirdrop,
  getAirdropList
} from '@/api/Airdrop/airdrop'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 导出组件
import ExportExcel from '@/components/exportExcel/exportExcel.vue'
// 导入组件
import ImportExcel from '@/components/exportExcel/importExcel.vue'
// 导出模板组件
import ExportTemplate from '@/components/exportExcel/exportTemplate.vue'


defineOptions({
    name: 'Airdrop'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const PublicChainOptions = ref([])
const formData = ref({
            airdropName: '',
            airdropPicture: "",
            airdropValue: 0,
            airdropEndtime: new Date(),
            publicChainId: '',
            airdropIsPass: false,
            airdropIsShow: false,
            airdropUrl: '',
            brief: '',
            detail: '',
            xLink: '',
            tgLink: '',
            discordLink: '',
            distributionDate: new Date(),
        })



// 验证规则
const rule = reactive({
               airdropName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               airdropPicture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               airdropValue : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               airdropEndtime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               distributionDate : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               publicChainId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               airdropUrl : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
        airdropEndtime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startAirdropEndtime && !searchInfo.value.endAirdropEndtime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startAirdropEndtime && searchInfo.value.endAirdropEndtime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startAirdropEndtime && searchInfo.value.endAirdropEndtime && (searchInfo.value.startAirdropEndtime.getTime() === searchInfo.value.endAirdropEndtime.getTime() || searchInfo.value.startAirdropEndtime.getTime() > searchInfo.value.endAirdropEndtime.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    if (searchInfo.value.airdropIsPass === ""){
        searchInfo.value.airdropIsPass=null
    }
    if (searchInfo.value.airdropIsShow === ""){
        searchInfo.value.airdropIsShow=null
    }
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getAirdropList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    PublicChainOptions.value = await getDictFunc('PublicChain')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteAirdropFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteAirdropByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateAirdropFunc = async(row) => {
    const res = await findAirdrop({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteAirdropFunc = async (row) => {
    const res = await deleteAirdrop({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        airdropName: '',
        airdropPicture: "",
        airdropValue: 0,
        airdropEndtime: new Date(),
        distributionDateL:new Date(),
        publicChainId: '',
        airdropIsPass: false,
        airdropIsShow: false,
        airdropUrl: '',
        brief: '',
        detail: '',
        xLink: '',
        tgLink: '',
        discordLink: '',

        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createAirdrop(formData.value)
                  break
                case 'update':
                  res = await updateAirdrop(formData.value)
                  break
                default:
                  res = await createAirdrop(formData.value)
                  break
              }
              btnLoading.value = false
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findAirdrop({ ID: row.ID })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>

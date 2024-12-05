
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
      
        <el-form-item label="预售项目名称" prop="presaleName">
         <el-input v-model="searchInfo.presaleName" placeholder="搜索条件" />
        </el-form-item>
        <el-form-item label="预售开始时间" prop="presaleStartTime">
            
            <template #label>
            <span>
              预售开始时间
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
            <el-date-picker v-model="searchInfo.startPresaleStartTime" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endPresaleStartTime ? time.getTime() > searchInfo.endPresaleStartTime.getTime() : false"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endPresaleStartTime" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startPresaleStartTime ? time.getTime() < searchInfo.startPresaleStartTime.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item label="上传者用户ID" prop="userId">
            
             <el-input v-model.number="searchInfo.userId" placeholder="搜索条件" />
        </el-form-item>
            <el-form-item label="预售项目是否审核通过" prop="presaleIsPass">
            <el-select v-model="searchInfo.presaleIsPass" clearable placeholder="请选择">
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
            <el-form-item label="预售项目是否展示" prop="presaleIsShow">
            <el-select v-model="searchInfo.presaleIsShow" clearable placeholder="请选择">
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
           <el-form-item label="公链ID" prop="publicChainId">
            <el-select v-model="searchInfo.publicChainId" clearable placeholder="请选择" @clear="()=>{searchInfo.publicChainId=undefined}">
              <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
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
            <ExportTemplate  template-id="Presale_Presale" />
            <ExportExcel  template-id="Presale_Presale" />
            <ImportExcel  template-id="Presale_Presale" @on-success="getTableData" />
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
        
          <el-table-column align="left" label="预售项目名称" prop="presaleName" width="120" />
          <el-table-column label="预售项目图片" prop="presalePicture" width="200">
              <template #default="scope">
                <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.presalePicture)" fit="cover"/>
              </template>
          </el-table-column>
         <el-table-column align="left" label="预售开始时间" prop="presaleStartTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.presaleStartTime) }}</template>
         </el-table-column>
          <el-table-column align="left" label="上传者用户ID" prop="userId" width="120" />
        <el-table-column align="left" label="预售项目是否审核通过" prop="presaleIsPass" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.presaleIsPass) }}</template>
        </el-table-column>
        <el-table-column align="left" label="预售项目是否展示" prop="presaleIsShow" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.presaleIsShow) }}</template>
        </el-table-column>
        <el-table-column align="left" label="公链ID" prop="publicChainId" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.publicChainId,PublicChainOptions) }}
            </template>
        </el-table-column>
          <el-table-column align="left" label="预售项目链接" prop="presaleurl" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updatePresaleFunc(scope.row)">编辑</el-button>
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
            <el-form-item label="预售项目名称:"  prop="presaleName" >
              <el-input v-model="formData.presaleName" :clearable="true"  placeholder="请输入预售项目名称" />
            </el-form-item>
            <el-form-item label="预售项目图片:"  prop="presalePicture" >
                <SelectImage
                 v-model="formData.presalePicture"
                 file-type="image"
                />
            </el-form-item>
            <el-form-item label="预售开始时间:"  prop="presaleStartTime" >
              <el-date-picker v-model="formData.presaleStartTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="上传者用户ID:"  prop="userId" >
              <el-input v-model.number="formData.userId" :clearable="true" placeholder="请输入上传者用户ID" />
            </el-form-item>
            <el-form-item label="预售项目是否审核通过:"  prop="presaleIsPass" >
              <el-switch v-model="formData.presaleIsPass" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="预售项目是否展示:"  prop="presaleIsShow" >
              <el-switch v-model="formData.presaleIsShow" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
            </el-form-item>
            <el-form-item label="公链ID:"  prop="publicChainId" >
              <el-select v-model="formData.publicChainId" placeholder="请选择公链ID" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in PublicChainOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="预售项目链接:"  prop="presaleurl" >
              <el-input v-model="formData.presaleurl" :clearable="true"  placeholder="请输入预售项目链接" />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close size="800" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="预售项目名称">
                        {{ detailFrom.presaleName }}
                    </el-descriptions-item>
                    <el-descriptions-item label="预售项目图片">
                            <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailFrom.presalePicture)" :src="getUrl(detailFrom.presalePicture)" fit="cover" />
                    </el-descriptions-item>
                    <el-descriptions-item label="预售开始时间">
                        {{ detailFrom.presaleStartTime }}
                    </el-descriptions-item>
                    <el-descriptions-item label="上传者用户ID">
                        {{ detailFrom.userId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="预售项目是否审核通过">
                        {{ detailFrom.presaleIsPass }}
                    </el-descriptions-item>
                    <el-descriptions-item label="预售项目是否展示">
                        {{ detailFrom.presaleIsShow }}
                    </el-descriptions-item>
                    <el-descriptions-item label="公链ID">
                        {{ detailFrom.publicChainId }}
                    </el-descriptions-item>
                    <el-descriptions-item label="预售项目链接">
                        {{ detailFrom.presaleurl }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createPresale,
  deletePresale,
  deletePresaleByIds,
  updatePresale,
  findPresale,
  getPresaleList
} from '@/api/Presale/presale'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'

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
    name: 'Presale'
})

// 提交按钮loading
const btnLoading = ref(false)

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const PublicChainOptions = ref([])
const formData = ref({
            presaleName: '',
            presalePicture: "",
            presaleStartTime: new Date(),
            userId: undefined,
            presaleIsPass: false,
            presaleIsShow: false,
            publicChainId: '',
            presaleurl: '',
        })



// 验证规则
const rule = reactive({
               presaleName : [{
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
               presalePicture : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               presaleStartTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               presaleIsPass : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               presaleIsShow : [{
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
               presaleurl : [{
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
        presaleStartTime : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startPresaleStartTime && !searchInfo.value.endPresaleStartTime) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startPresaleStartTime && searchInfo.value.endPresaleStartTime) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startPresaleStartTime && searchInfo.value.endPresaleStartTime && (searchInfo.value.startPresaleStartTime.getTime() === searchInfo.value.endPresaleStartTime.getTime() || searchInfo.value.startPresaleStartTime.getTime() > searchInfo.value.endPresaleStartTime.getTime())) {
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
    if (searchInfo.value.presaleIsPass === ""){
        searchInfo.value.presaleIsPass=null
    }
    if (searchInfo.value.presaleIsShow === ""){
        searchInfo.value.presaleIsShow=null
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
  const table = await getPresaleList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deletePresaleFunc(row)
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
      const res = await deletePresaleByIds({ IDs })
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
const updatePresaleFunc = async(row) => {
    const res = await findPresale({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePresaleFunc = async (row) => {
    const res = await deletePresale({ ID: row.ID })
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
        presaleName: '',
        presalePicture: "",
        presaleStartTime: new Date(),
        userId: undefined,
        presaleIsPass: false,
        presaleIsShow: false,
        publicChainId: '',
        presaleurl: '',
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
                  res = await createPresale(formData.value)
                  break
                case 'update':
                  res = await updatePresale(formData.value)
                  break
                default:
                  res = await createPresale(formData.value)
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
  const res = await findPresale({ ID: row.ID })
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

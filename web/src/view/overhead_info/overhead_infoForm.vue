<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="创建日期:" prop="day">
          <el-date-picker v-model="formData.day" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="分类id:" prop="category_id">
          <el-input v-model.number="formData.category_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="消费金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="描述:" prop="desc">
          <el-input v-model="formData.desc" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'OverheadInfo'
}
</script>

<script setup>
import {
  createOverheadInfo,
  updateOverheadInfo,
  findOverheadInfo
} from '@/api/overhead_info'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            uid: 0,
            day: new Date(),
            category_id: 0,
            name: '',
            amount: 0,
            desc: '',
        })
// 验证规则
const rule = reactive({
               uid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               day : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               category_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               amount : [{
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
      const res = await findOverheadInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createOverheadInfo(formData.value)
               break
             case 'update':
               res = await updateOverheadInfo(formData.value)
               break
             default:
               res = await createOverheadInfo(formData.value)
               break
           }
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

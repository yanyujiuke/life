<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户id:" prop="uid">
          <el-input v-model.number="formData.uid" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="月份:" prop="month">
          <el-date-picker v-model="formData.month" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="每月开销金额:" prop="amount">
          <el-input-number v-model="formData.amount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="每月开销单数:" prop="num">
          <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="每月消费金额排名:" prop="rank">
          <el-input v-model.number="formData.rank" :clearable="true" placeholder="请输入" />
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
  name: 'OverheadMonthStatistics'
}
</script>

<script setup>
import {
  createOverheadMonthStatistics,
  updateOverheadMonthStatistics,
  findOverheadMonthStatistics
} from '@/api/overhead_month_statistics'

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
            month: new Date(),
            amount: 0,
            num: 0,
            rank: 0,
            desc: '',
        })
// 验证规则
const rule = reactive({
               uid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               month : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               amount : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               num : [{
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
      const res = await findOverheadMonthStatistics({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reMonthStatistics
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
               res = await createOverheadMonthStatistics(formData.value)
               break
             case 'update':
               res = await updateOverheadMonthStatistics(formData.value)
               break
             default:
               res = await createOverheadMonthStatistics(formData.value)
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

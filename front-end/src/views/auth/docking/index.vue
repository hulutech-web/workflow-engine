<script setup lang="ts">
import type {StepsProps} from 'naive-ui'

const currentStatus = ref<StepsProps['status']>('process')
const current = ref(1)
const prev = () => {
  if (current.value === 1) {
    return
  }
  current.value--
}
const next = () => {
  if (current.value === 5) {
    return
  }
  current.value++
}

const submitForm = ref({
  tenant_name: "",
  key_id: "",
  key_value: "",
  smsapi:"",
  smstemplate:"",
  wxapi:"",
  wxtemplate:""
})
const onsubmit = () => {

}
</script>

<template>
  <div>
    <n-card title="接入配置" size="small">

      <n-space vertical>
        <n-steps :current="current as number" :status="currentStatus">
          <n-step
            title="租户"
            description="不同租户彼此隔离，安全有保障"
          />
          <n-step
            title="部门"
            description="匹配部门信息，和部门经理，部门主管"
          />
          <n-step
            title="员工"
            description="匹配后可以发起流程和审批流程，并接收通知"
          />
          <n-step
            title="消息通知"
            description="有效的消息通知，配置系统参数，包括短信和微信消息通知API参数"
          />

          <n-step
            title="使用指南"
            description="根据指南完成流程设置"
          />
        </n-steps>
      </n-space>
      <div class="min-h-[500px] flex justify-center items-center">
        <n-form style="width:600px;" label-placement="left">
          <div v-show="current==1">
            <n-form-item label="租户名">
              <n-input v-model:value="submitForm.tenant_name" placeholder="请输入流程名称"/>
            </n-form-item>
            <n-form-item label="key_id">
              <n-input v-model:value="submitForm.key_id" placeholder="请输入流程名称"/>
            </n-form-item>
            <n-form-item label="密码">
              <n-input v-model:value="submitForm.key_value" type="password" placeholder="请输入流程名称"/>
            </n-form-item>
          </div>
          <div v-show="current==2">
            <a href="" class="text-blue-500">模板下载</a>
            <n-form-item label="导入部门">
              <n-upload
                multiple
                directory-dnd
                action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
                :max="5"
              >
                <n-upload-dragger>
                  <div style="margin-bottom: 12px">
                    <n-icon size="48" :depth="3">
                    </n-icon>
                  </div>
                  <n-text style="font-size: 16px">
                    点击或者拖动文件到该区域来上传
                  </n-text>
                  <n-p depth="3" style="margin: 8px 0 0 0">
                    请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码
                  </n-p>
                </n-upload-dragger>
              </n-upload>
            </n-form-item>
          </div>
          <div v-show="current==3">
            <a href="" class="text-blue-500">模板下载</a>
            <n-form-item label="导入员工">
              <n-upload
                multiple
                directory-dnd
                action="https://www.mocky.io/v2/5e4bafc63100007100d8b70f"
                :max="5"
              >
                <n-upload-dragger>
                  <div style="margin-bottom: 12px">
                    <n-icon size="48" :depth="3">
                    </n-icon>
                  </div>
                  <n-text style="font-size: 16px">
                    点击或者拖动文件到该区域来上传
                  </n-text>
                  <n-p depth="3" style="margin: 8px 0 0 0">
                    请不要上传敏感数据，比如你的银行卡号和密码，信用卡号有效期和安全码
                  </n-p>
                </n-upload-dragger>
              </n-upload>
            </n-form-item>
          </div>
          <div v-show="current==4">
            <n-form-item label="短信api">
              <n-input v-model:value="submitForm.smsapi" placeholder="请输入短信api"></n-input>
            </n-form-item>
            <n-form-item label="短信模板">
              <n-input v-model:value="submitForm.smstemplate" placeholder="请输入短信模板"></n-input>
            </n-form-item>
            <n-form-item label="微信推送API">
              <n-input v-model:value="submitForm.wxapi" placeholder="请输入微信推送API"></n-input>
            </n-form-item>
            <n-form-item label="微信推送模板">
              <n-input v-model:value="submitForm.wxtemplate" placeholder="请输入微信推送模板"></n-input>
            </n-form-item>
          </div>

          <div v-show="current==5">
            <div class="b-1 p-3 mb-3  b-dotted b-red">
              使用指南：
              <a href="https://hulutech-web.github.io/goravel-workflow.github.io/guides/hooks.html" class="text-blue-500">详情</a><br/>
              介绍<br/>
              采用go的组合方式，实现2个接口即可，即可获得框架提供的能力，开发者可自行进行实现后续逻辑，如发送审核成功后通知消息，审核失败后的后续处理等等。<br/>
              接口1-通知发起人：NotifySendOne<br/>
              接口2-通知下一审批人：NotifyNextAuditor<br/>
              <br/>
              集成于使用<br/>
              模型关联<br/>
              定义关联模型，本案以User结构为例。<br/>
              定义结构体，注入服务,WorkNo,Password为必须<br/>
              定义接口，实现2个接口<br/>
              实例化workflow<br/>
              注册workflow，在app/providers/app_services_provider.go中实例化workflow，并注入hooks<br/>
              实现效果 NotifySendOne<br/>
              当流程执行过程中，流程被驳回时，将自动调用NotifySendOne方法，传递的id参数为emp.id，表示通知发起人。后续由开发者自行实现逻辑，例如：发送邮件，发送短信，消息推送等。<br/>
              当流程执行过程中，整个流程执行完毕时（所有人都同意），将自动调用NotifySendOne方法，传递的id参数为emp.id，表示通知发起人。后续由开发者自行实现逻辑，例如：发送邮件，发送短信，消息推送等。<br/>
              NotifyNextAuditor 当流程执行过程中，其中某一个环节审批通过时，将自动调用NotifyNextAuditor方法，传递的id参数为emp.id，表示通知下一个审批人。后续由开发者自行实现逻辑，例如：发送邮件，发送短信，消息推送等。
              <br/>
            </div>

          </div>

        </n-form>
      </div>

      <div class="flex justify-center items-center">
        <n-space>
          <n-button @click="prev" type="primary">
            上一步
          </n-button>
          <n-button @click="next" type="error">
            下一步
          </n-button>
        </n-space>
      </div>

      <div class="flex justify-center items-center mt-2">
        <n-button @click="onsubmit" type="primary" size="large">
          提交
        </n-button>
      </div>

    </n-card>
  </div>
</template>

<style scoped>

</style>

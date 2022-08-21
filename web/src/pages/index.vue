<script setup lang="ts">
import type { Paste } from '~/composables/types'

const types = Array.from(languages.keys())

const expired_days_dict = {
  '1 Day': 1,
  '1 Week': 7,
  '1 Month': 30,
}

const expired_days = $ref('1 Week')

const paste = $ref({ expired_days: 7, type: 'Text' } as Paste)

function setCode(code: string) {
  paste.data = code
}

const router = useRouter()

function onPaste() {
  fetch(`${import.meta.env.VITE_BASE_URL}/create`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(paste),
  })
    .then(res => res.json(),
    )
    .then((res) => {
      router.push(`/${res.uuid}`)
    })
    .catch((err) => {
      console.error(err)
    })
}

function onChange() {
  paste.expired_days = expired_days_dict[expired_days as keyof typeof expired_days_dict]
}
</script>

<template>
  <div flex="~" flex-col items-center justify-center>
    <div flex="~" flex-row items-center justify-center space-x-5 pb-5>
      <span>
        Type:
        <select v-model="paste.type">
          <option v-for="type in types" :key="type" :value="type">{{ type }}</option>
        </select>
      </span>
      <span>
        Expired Days:
        <select v-model="expired_days" @change="onChange">
          <option v-for="x in Object.keys(expired_days_dict)" :key="x">{{ x }}</option>
        </select>
      </span>
      <button btn @click="onPaste">
        paste
      </button>
    </div>
    <CodeEditor :paste="paste" @set-code="setCode" />
  </div>
</template>

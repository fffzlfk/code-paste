<script setup lang="ts">
import type { Paste } from '~/composables/types'

const { id } = defineProps<{ id: string }>()

let paste = $ref({} as Paste)

const router = useRouter()

function fetchPaste(id: string) {
  fetch(`${import.meta.env.VITE_BASE_URL}/read/${id}`)
    .then(res => res.json())
    .then((res) => {
      paste = res
    })
    .catch((err) => {
      console.error(err)
      router.push('/')
    })
}

onMounted(() => {
  fetchPaste(id)
})
</script>

<template>
  <div flex="~" flex-col items-center justify-center>
    <div flex="~" flex-row items-center justify-center space-x-5 pb-5>
      <span>Type:
        <select disabled>
          <option value="paste.type">{{ paste.type }}</option>
        </select>
      </span>
      <span>ExpiredAt:
        {{ UTC2Local(paste.expired_at) }}
      </span>
    </div>
    <CodeEditor :paste="paste" />
  </div>
</template>

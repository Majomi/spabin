<template>
  <div class="flex flex-col">
    <div class="card-title">{{ title }}</div>
    <div class="card-content">{{ content }}</div>
  </div>
</template>

<script>
export default {
  props: {
    title: {
      type: String,
      required: true
    },
    content: {
      type: String,
      required: true
    }
  }
};
</script>

<style scoped></style>

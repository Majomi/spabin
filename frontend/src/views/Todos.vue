<template>
  <div>
    <Todo
      v-for="(todo, index) in todos"
      :key="index"
      :title="todo.title"
      :content="todo.msg"
    />
  </div>
</template>

<script>
import Todo from "@/components/Todo.vue";
export default {
  components: {
    Todo
  },
  data() {
    return {
      todos: []
    };
  },
  async created() {
    await this.fetchData();
  },
  methods: {
    async fetchData() {
      console.log("fetching data");

      const { data } = await this.$axios.get("/api/todos");
      this.todos = data;
    }
  }
};
</script>

<style scoped></style>

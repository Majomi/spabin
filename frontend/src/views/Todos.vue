<template>
  <div>
    <Todo
      v-for="todo in todos"
      :key="todo.title"
      :title="todo.title"
      :content="todo.content"
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
      todos: [
        {
          title: "Title",
          content: "Content"
        }
      ]
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

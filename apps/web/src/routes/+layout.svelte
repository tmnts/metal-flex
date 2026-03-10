<script>
  import '../app.css'; // <--- ВОТ ОН, ГЛАВНЫЙ ИМПОРТ!
  let { children } = $props();
</script>

{@render children()}

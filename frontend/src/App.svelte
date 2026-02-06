<script>
  import Router from "svelte-spa-router";
  import AdminLayout from "./layouts/AdminLayout.svelte";
  import DocumentList from "./pages/Documents/DocumentList.svelte";
  import DocumentAdd from "./pages/Documents/DocumentAdd.svelte";
  import DocumentEdit from "./pages/Documents/DocumentEdit.svelte";
  import NotFound from "./pages/NotFound.svelte";

  // Public routes that don't need AdminLayout (id 404)
  const routes = {
    "/": DocumentList,
    "/documents": DocumentList,
    "/documents/add": DocumentAdd,
    "/documents/edit/:id": DocumentEdit,
    "*": NotFound,
  };
</script>

<AdminLayout>
  <Router {routes} />
</AdminLayout>

<script>
  import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";

  const queryClient = new QueryClient();

  import Login from "./pages/Login.svelte";
  import Register from "./pages/Register.svelte";
  import Home from "./pages/Home.svelte";
  import Dashboard from "./pages/Dashboard.svelte";
  import { Router, Route } from "svelte-routing";
  import Project from "./pages/Project.svelte";
  import Task from "./pages/Task.svelte";
  import Setting from "./pages/Setting.svelte";
  import Toaster from "./components/Toaster.svelte";
  import { setToastState } from "./lib/toast";
  import Install from "./pages/Install.svelte";
  export let url = "";
  setToastState();
</script>

<Toaster />

<QueryClientProvider client={queryClient}>
  <Router {url}>
    <Route path="/" component={Home} />
    <Route path="/login" component={Login} />
    <Route path="/register" component={Register} />
    <Route path="/dashboard" component={Dashboard} />
    <Route path="/projects" component={Project} />
    <Route path="/tasks" component={Task} />
    <Route path="/settings" component={Setting} />
    <Route path="/install" component={Install} />
  </Router>
</QueryClientProvider>

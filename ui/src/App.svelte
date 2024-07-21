<script>
  import { QueryClient, QueryClientProvider } from "@tanstack/svelte-query";
  import { Router, Route } from "svelte-routing";

  import Login from "./pages/Login.svelte";
  import Register from "./pages/Register.svelte";
  import Home from "./pages/Home.svelte";
  import Dashboard from "./pages/Dashboard.svelte";
  import Project from "./pages/Project.svelte";
  import Config from "./pages/Config.svelte";
  import Install from "./pages/Install.svelte";
  import Toaster from "./components/Toaster.svelte";
  import { setToastState } from "./lib/toast";
  import Ssl from "./pages/Ssl.svelte";

  export let url = "";
  const queryClient = new QueryClient();
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
    <Route path="/config" component={Config} />
    <Route path="/ssl" component={Ssl} />
    <Route path="/install" component={Install} />
  </Router>
</QueryClientProvider>

<script lang="ts">
  import Button from "../components/ui/Button.svelte";
  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import PlusIcon from "lucide-svelte/icons/plus";
  import BaseLayout from "../components/BaseLayout.svelte";

  // Sample data for SSL configurations
  let sslConfigs = [
    {
      id: 1,
      domain: "example.com",
      directory: "/var/www/example",
      status: "Active",
    },
    {
      id: 2,
      domain: "test.org",
      directory: "/var/www/test",
      status: "Expired",
    },
    {
      id: 3,
      domain: "myapp.io",
      directory: "/var/www/myapp",
      status: "Renewing",
    },
    {
      id: 4,
      domain: "blog.dev",
      directory: "/var/www/blog",
      status: "Inactive",
    },
  ];

  // Function to determine status color
  function getStatusColor(status: string) {
    switch (status.toLowerCase()) {
      case "active":
        return "text-green-600";
      case "renewing":
        return "text-blue-600";
      case "expired":
        return "text-red-600";
      case "inactive":
        return "text-yellow-600";
      default:
        return "text-gray-600";
    }
  }
</script>

<BaseLayout>
  <div
    class="bg-pattern absolute inset-x-0 top-0 -z-0 h-[280px] border-b border-gray-500"
  ></div>

  <div class="container mx-auto max-w-6xl px-4 pb-12 relative min-h-screen">
    <div class="py-12 pb-4 flex justify-between items-center">
      <h1 class="text-4xl font-bold">SSL Configurations</h1>
      <Button><PlusIcon class="mr-1 w-4 h-4" /> Add SSL Configuration</Button>
    </div>

    <div class="py-8">
      <RetroOutline className="w-full" childClassName="p-0">
        <table class="ssl-table table-auto">
          <thead class="border-b border-gray-800">
            <tr>
              <th>ID</th>
              <th>Domain</th>
              <th>Directory</th>
              <th>Status</th>
              <th class="w-64">Actions</th>
            </tr>
          </thead>
          <tbody>
            {#each sslConfigs as config}
              <tr>
                <td>{config.id}</td>
                <td>{config.domain}</td>
                <td>{config.directory}</td>
                <td class={getStatusColor(config.status)}>{config.status}</td>
                <td>
                  <Button className="mr-2" variant="outline-primary"
                    >View</Button
                  >
                  <Button className="mr-2" variant="warning">Renew</Button>
                  <Button variant="danger">Delete</Button>
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </RetroOutline>
    </div>
  </div>
</BaseLayout>

<style>
  .ssl-table {
    width: 100%;
    border-collapse: collapse;
  }
  .ssl-table th,
  .ssl-table td {
    padding: 12px;
    border: none;
    text-align: left;
  }
</style>

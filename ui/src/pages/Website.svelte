<script>
  import Button from "../components/ui/Button.svelte";
  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import PlusIcon from "lucide-svelte/icons/plus";
  import BaseLayout from "../components/BaseLayout.svelte";
  import Dialog from "../components/ui/Dialog.svelte";

  let websites = [
    {
      id: 1,
      domain: "example.com",
      folder: "/var/www/example",
      status: "Active",
    },
    { id: 2, domain: "test.com", folder: "/var/www/test", status: "Inactive" },
    {
      id: 3,
      domain: "demo.net",
      folder: "/var/www/demo",
      status: "Maintenance",
    },
  ];

  let showDialog = false;

  let newWebsite = {
    domain: "",
    folder: "",
    status: "Active",
  };

  function toggleDialog() {
    showDialog = !showDialog;
  }

  function addWebsite() {
    websites = [...websites, { id: websites.length + 1, ...newWebsite }];
    newWebsite = { domain: "", folder: "", status: "Active" };
    toggleDialog();
  }
</script>

<BaseLayout>
  <div
    class="bg-pattern absolute inset-x-0 top-0 -z-0 h-[280px] border-b border-gray-500"
  ></div>

  <div class="container mx-auto max-w-6xl px-4 pb-12 relative min-h-screen">
    <div class="py-12 pb-4 flex justify-between items-center">
      <h1 class="text-4xl font-bold">Websites</h1>
      <Button
        on:click={toggleDialog}
        className="inline-flex items-center gap-2 pl-3"
      >
        <PlusIcon class="size-4" /> Add Website</Button
      >
    </div>

    <div class="py-8">
      <RetroOutline className="w-full">
        <table class="web-table w-full">
          <thead class="border-b border-gray-800">
            <tr class="text-left">
              <th class="px-4 w-14">ID</th>
              <th class="p-2">Domain</th>
              <th class="p-2">Folder</th>
              <th class="p-2">Status</th>
              <th class="p-2 w-56">Actions</th>
            </tr>
          </thead>
          <tbody>
            {#each websites as website}
              <tr>
                <td class="p-2 px-4">{website.id}</td>
                <td class="p-2">{website.domain}</td>
                <td class="p-2">{website.folder}</td>
                <td class="p-2">
                  <span
                    class="px-2 py-1 rounded-full {website.status ===
                    'Active'
                      ? 'text-green-500'
                      : website.status === 'Inactive'
                        ? 'text-red-500'
                        : 'text-yellow-500'}"
                  >
                    {website.status}
                  </span>
                </td>
                <td class="p-2">
                  <Button className="mr-2" variant="outline-primary">View</Button>
                  <Button className="mr-2" variant="warning" >Edit</Button>
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

<Dialog bind:open={showDialog} on:close={toggleDialog} title="Add new website">
  <form on:submit|preventDefault={addWebsite} class="space-y-4">
    <div>
      <label for="domain" class="block text-sm font-medium text-gray-700"
        >Domain</label
      >
      <input
        type="text"
        id="domain"
        bind:value={newWebsite.domain}
        required
        class="w-full px-3 py-2 border border-gray-300 mt-1"
      />
    </div>
    <div>
      <label for="folder" class="block text-sm font-medium text-gray-700"
        >Folder</label
      >
      <input
        type="text"
        id="folder"
        bind:value={newWebsite.folder}
        required
        class="w-full px-3 py-2 border border-gray-300 mt-1"
      />
    </div>
    <div>
      <label for="status" class="block text-sm font-medium text-gray-700"
        >Status</label
      >
      <select
        id="status"
        bind:value={newWebsite.status}
        class="w-full px-3 py-2 border border-gray-300 mt-1"
      >
        <option value="Active">Active</option>
        <option value="Inactive">Inactive</option>
        <option value="Maintenance">Maintenance</option>
      </select>
    </div>
    <div class="flex justify-end space-x-2">
      <Button type="button" on:click={toggleDialog}>Cancel</Button>
      <Button type="submit">Submit</Button>
    </div>
  </form>
</Dialog>

<style>
    .web-table {
      width: 100%;
      border-collapse: collapse;
    }
    .web-table th,
    .web-table td {
      padding: 12px;
      border: none;
      text-align: left;
    }
  </style>
  
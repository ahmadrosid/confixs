<script lang="ts">
  import Button from "../components/ui/Button.svelte";
  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import BaseLayout from "../components/BaseLayout.svelte";
  import { pb } from "../lib/pocketbase";

  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";

  const defaultUserPreferences = {
    darkMode: false,
    emailNotifications: true,
    pushNotifications: false,
  };

  let user = {
    name: "",
    username: "",
    email: "",
    password: "",
    new_password: "",
    new_confirm_password: "",
    photoUrl: "",
    ...defaultUserPreferences,
  };

  let activeSection = "account";

  onMount(async () => {
    if (!pb.authStore.isValid) {
      navigate("/login");
      return;
    }

    const authUser = pb.authStore.model;
    if (authUser) {
      user = {
        ...user,
        username: authUser.username,
        email: authUser.email,
        photoUrl: authUser.avatar,
        name: authUser.name,
        // In a real app, fetch these from your backend
        ...defaultUserPreferences,
      };
    }
  });

  const saveSettings = () => {
    // TODO: implement save settings
  };

  const setActiveSection = (section: string) => {
    activeSection = section;
  };
</script>

<BaseLayout>
  <div
    class="bg-pattern absolute inset-x-0 top-0 -z-0 h-[280px] border-b border-gray-500"
  ></div>

  <div class="container mx-auto max-w-6xl px-4 pb-12 relative">
    <div class="py-12 pb-4">
      <h1 class="text-4xl font-bold">Settings</h1>
    </div>

    <div class="flex gap-8 py-8">
      <!-- Sidebar -->
      <RetroOutline className="max-w-64 h-fit">
        <nav>
          <ul>
            <li>
              <button
                class="sidebar-link"
                class:active={activeSection === "account"}
                on:click={() => setActiveSection("account")}
              >
                Account
              </button>
            </li>
            <li>
              <button
                class="sidebar-link"
                class:active={activeSection === "preferences"}
                on:click={() => setActiveSection("preferences")}
              >
                Preferences
              </button>
            </li>
            <li>
              <button
                class="sidebar-link"
                class:active={activeSection === "notifications"}
                on:click={() => setActiveSection("notifications")}
              >
                Notifications
              </button>
            </li>
          </ul>
        </nav>
      </RetroOutline>

      <!-- Main content -->
      <div class="flex-1">
        <form on:submit|preventDefault={saveSettings}>
          {#if activeSection === "account"}
            <RetroOutline className="mb-8">
                <div class="p-6">
                    <h2 class="text-2xl font-bold mb-4">Account</h2>
                    <div class="space-y-4">
                        <div>
                        <label
                            for="name"
                            class="block text-sm font-medium text-gray-700"
                            >Full Name</label
                        >
                        <input
                            type="text"
                            id="name"
                            bind:value={user.name}
                            class="mt-1 block w-full border border-gray-800 shadow-sm p-2"
                        />
                        </div>
                        <div>
                        <label
                            for="username"
                            class="block text-sm font-medium text-gray-700"
                            >Full name</label
                        >
                        <input
                            type="text"
                            id="username"
                            bind:value={user.username}
                            class="mt-1 block w-full border border-gray-800 shadow-sm p-2"
                        />
                        </div>
                        <div>
                        <label
                            for="email"
                            class="block text-sm font-medium text-gray-700">Email</label
                        >
                        <input
                            type="email"
                            id="email"
                            bind:value={user.email}
                            class="mt-1 block w-full border border-gray-800 shadow-sm p-2"
                        />
                        </div>
                    </div>
                </div>
            </RetroOutline>

            <RetroOutline className="mb-8">
                <div class="p-6">
                    <h2 class="text-2xl font-bold mb-4">Change password</h2>
                    <div class="space-y-4">
                        <div>
                        <label
                            for="password"
                            class="block text-sm font-medium text-gray-700"
                            >Old Password</label
                        >
                        <input
                            type="password"
                            id="password"
                            bind:value={user.password}
                            class="mt-1 block w-full border border-gray-800 shadow-sm p-2"
                        />
                        </div>
                        <div>
                        <label
                            for="new_password"
                            class="block text-sm font-medium text-gray-700"
                            >New Password</label
                        >
                        <input
                            type="password"
                            id="new_password"
                            bind:value={user.new_password}
                            class="mt-1 block w-full border border-gray-800 shadow-sm p-2"
                        />
                        </div>
                        <div>
                        <label
                            for="new_confirm_password"
                            class="block text-sm font-medium text-gray-700"
                            >Confirm New Password</label
                        >
                        <input
                            type="password"
                            id="new_confirm_password"
                            bind:value={user.new_confirm_password}
                            class="mt-1 block w-full border border-gray-800 shadow-sm p-2"
                        />
                        </div>
                    </div>
                </div>
            </RetroOutline>
          {/if}

          {#if activeSection === "preferences"}
            <RetroOutline className="mb-8">
                <div class="p-6">
                    <h2 class="text-2xl font-bold mb-4">Preferences</h2>
                    <div class="space-y-4">
                        <div class="flex items-center">
                        <input
                            type="checkbox"
                            id="darkMode"
                            bind:checked={user.darkMode}
                            class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                        />
                        <label for="darkMode" class="ml-2 block text-sm text-gray-900"
                            >Dark Mode</label
                        >
                        </div>
                    </div>
                </div>
            </RetroOutline>
          {/if}

          {#if activeSection === "notifications"}
            <RetroOutline className="mb-8">
                <div class="p-6">

                    <h2 class="text-2xl font-bold mb-4">Notifications</h2>
                    <div class="space-y-4">
                        <div class="flex items-center">
                        <input
                            type="checkbox"
                            id="emailNotifications"
                            bind:checked={user.emailNotifications}
                            class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                        />
                        <label
                            for="emailNotifications"
                            class="ml-2 block text-sm text-gray-900"
                            >Email Notifications</label
                        >
                        </div>
                        <div class="flex items-center">
                        <input
                            type="checkbox"
                            id="pushNotifications"
                            bind:checked={user.pushNotifications}
                            class="h-4 w-4 text-indigo-600 focus:ring-indigo-500 border-gray-300 rounded"
                        />
                        <label
                            for="pushNotifications"
                            class="ml-2 block text-sm text-gray-900"
                            >Push Notifications</label
                        >
                        </div>
                    </div>
                </div>
            </RetroOutline>
          {/if}

          <div class="flex justify-end">
            <Button size="lg">Save Settings</Button>
          </div>
        </form>
      </div>
    </div>
  </div>
</BaseLayout>

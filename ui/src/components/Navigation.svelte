<script lang="ts">
  import { Link } from "svelte-routing";
  import Button from "./ui/Button.svelte";
  import { onMount } from "svelte";
  import { navigate } from "svelte-routing";
  import { pb, type UserModel } from "../lib/pocketbase";
  import RetroOutline from "./ui/RetroOutline.svelte";
  import { fade } from "svelte/transition";
  import UserSquare from "lucide-svelte/icons/square-user-round"

  let isLoggedIn = false;
  let user: UserModel;
  let isMenuOpen = false;
  let menuRef: HTMLDivElement;

  onMount(() => {
    checkAuthStatus();
    document.addEventListener("click", handleClickOutside);
    return () => {
      document.removeEventListener("click", handleClickOutside);
    };
  });

  function checkAuthStatus() {
    isLoggedIn = pb.authStore.isValid;
    if (isLoggedIn) {
      user = pb.authStore.model as UserModel;
    }
  }

  function handleLogout() {
    pb.authStore.clear();
    checkAuthStatus();
    navigate("/");
    isMenuOpen = false;
  }

  function handleClickOutside(event: MouseEvent) {
    if (menuRef && !event.target && !menuRef.contains(event.target)) {
      isMenuOpen = false;
    }
  }

  function toggleMenu() {
    isMenuOpen = !isMenuOpen;
  }
</script>

<nav class="relative z-10 border-b border-gray-800 bg-white">
  <div class="max-w-6xl mx-auto px-4">
    <div class="flex justify-between items-center py-4">
      <div class="flex gap-4 items-center">
        <Link to="/" class="font-semibold text-xl text-gray-950">
          Confixs
        </Link>
      </div>
      <div class="flex items-center space-x-8">
        {#if !isLoggedIn}
          <Link to="/login">
            <span class="text-gray-950 hover:text-gray-800 hover:underline"
              >Login</span
            >
          </Link>
          <Link to="/register">
            <Button>Register</Button>
          </Link>
        {:else}
          <div class="flex items-center space-x-8">
            <Link to="/dashboard" class="text-sm hover:underline">Dashboard</Link>

            <Link to="/projects" class="text-sm hover:underline">Websites</Link>

            <Link to="/tasks" class="text-sm hover:underline">SSL</Link>

            <Link to="/dashboard" class="text-sm hover:underline">Config</Link>

            <Link to="/settings" class="text-sm hover:underline">Settings</Link>
          </div>

          <div class="relative" bind:this={menuRef}>
            <button
              on:click={toggleMenu}
              class="flex items-center space-x-2 text-gray-950 hover:text-gray-800 focus:outline-none"
            >
              <UserSquare class="w-6 h-6" />
            </button>

            {#if isMenuOpen}
              <div
                class="absolute right-0 mt-3 w-48"
                transition:fade={{ delay: 30, duration: 50 }}
              >
                <RetroOutline>
                  <Link
                    to="/profile"
                    class="block px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    My Profile
                  </Link>
                  <button
                    on:click={handleLogout}
                    class="block w-full text-left px-4 py-2 text-sm text-gray-700 hover:bg-gray-100"
                  >
                    Logout
                  </button>
                </RetroOutline>
              </div>
            {/if}
          </div>
        {/if}
      </div>
    </div>
  </div>
</nav>

<script lang="ts">
  import BaseLayout from "../components/BaseLayout.svelte";
  import RetroOutline from "../components/ui/RetroOutline.svelte";
  import { Link, navigate } from "svelte-routing";
  import { pb } from '../lib/pocketbase';
  import EyeIcon from 'lucide-svelte/icons/eye'

  let email = '';
  let password = '';
  let error = '';

  async function handleLogin() {
    try {
      const authData = await pb.collection('users').authWithPassword(email, password);
      navigate('/dashboard');
    } catch (err) {
      console.error('Login failed:', err);
      error = 'Invalid email or password';
    }
  }

</script>

<BaseLayout>
  <div class="bg-white flex items-center justify-center min-h-screen">
    <RetroOutline className='max-w-md'>
      <div class="bg-white p-6">
        <h1 class="text-2xl font-bold mb-6">Login to your account</h1>

        <form on:submit|preventDefault={handleLogin}>
          <div class="mb-4">
            <label for="email" class="block mb-2"
              >Email or Username <span class="text-red-500">*</span></label
            >
            <p class="text-sm text-gray-500 mb-1">
              Either email or username is accepted.
            </p>
            <input
              type="text"
              id="email"
              name="email"
              class="w-full px-3 py-2 border border-gray-300"
              bind:value={email}
              required
            />
          </div>

          <div class="mb-6">
            <label for="password" class="block mb-2"
              >Password <span class="text-red-500">*</span></label
            >
            <div class="relative">
              <input
                type="password"
                id="password"
                bind:value={password}
                name="password"
                class="w-full px-3 py-2 border border-gray-300 pr-10"
                required
              />
              <span
                class="absolute inset-y-0 right-0 flex items-center pr-3 cursor-pointer"
              >
                <EyeIcon className='size-4 text-gray-500'/>
              </span>
            </div>
          </div>

          {#if error}
            <div class="text-red-500 text-sm pb-4">{error}</div>
          {/if}
          <button
            type="submit"
            class="w-full bg-black text-white py-2 hover:bg-gray-800"
            >Login</button
          >
        </form>

        <p class="mt-4 text-center">
          Don't have an account? <Link
            to="/register"
            class="text-blue-600 hover:underline">Register</Link
          >
        </p>
      </div>
    </RetroOutline>
  </div>
</BaseLayout>

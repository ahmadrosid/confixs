<script lang="ts">
  import BaseLayout from "../components/BaseLayout.svelte";
  import NginxInstall from "../components/NginxInstall.svelte";
  import RetroOutline from "../components/ui/RetroOutline.svelte";
</script>

<BaseLayout>
  <div class="bg-white flex items-start justify-center min-h-screen">
    <div class="pt-12">
        <RetroOutline className="w-full sm:w-[48rem] max-w-3xl">
          <div class="bg-white p-6 flex-1 w-full">
            <NginxInstall />
          </div>
        </RetroOutline>
    </div>
  </div>
</BaseLayout>

<script>
  import Shield from "lucide-svelte/icons/shield";
  import PanelsTopLeft from "lucide-svelte/icons/panels-top-left";
  import Terminal from "lucide-svelte/icons/terminal";
  import Server from "lucide-svelte/icons/server";

  import RetroOutline from "../ui/RetroOutline.svelte";

  let features = [
    {
      icon: PanelsTopLeft,
      title: "User-Friendly Interface",
      description:
        "Navigate complex Nginx setups with ease using our clean, intuitive design.",
    },
    {
      icon: Shield,
      title: "One-Click SSL Integration",
      description:
        "Secure your sites in seconds with automatic Let's Encrypt SSL certificate issuance and renewal",
    },
    {
      icon: Terminal,
      title: "Advanced Configuration Editor",
      description:
        "Fine-tune your Nginx settings with our built-in editor, complete with syntax highlighting and validation.",
    },
    {
      icon: Server,
      title: "Multi-Site Management",
      description:
        "Whether you're running one site or one hundred, Confixs scales to meet your needs.",
    },
  ];
</script>

<div class="bg-gray-50 py-16">
  <div class="container max-w-6xl mx-auto px-4">
    <h2 class="text-3xl font-bold mb-12">Features</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-8">
      {#each features as feature}
        <RetroOutline>
          <div class="p-6 bg-white">
            <div class="flex items-center mb-4">
              <div
                class="w-12 h-12 flex items-center justify-center rounded-full bg-green-200 mr-4"
              >
                <svelte:component
                  this={feature.icon}
                  class="w-6 h-6 text-green-600"
                />
              </div>
              <h3 class="text-xl font-semibold">{feature.title}</h3>
            </div>
            <p class="text-gray-600">{feature.description}</p>
          </div>
        </RetroOutline>
      {/each}
    </div>
  </div>
</div>

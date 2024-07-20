<script>
  import Building from "lucide-svelte/icons/building";
  import GraduationCap from "lucide-svelte/icons/graduation-cap";
  import Users from "lucide-svelte/icons/users";
  import Briefcase from "lucide-svelte/icons/briefcase";

  const useCases = [
    {
      icon: Building,
      title: "Enterprise Internship Programs",
      description:
        "Streamline large-scale internship management across multiple departments, ensuring consistent experiences and efficient oversight.",
    },
    {
      icon: GraduationCap,
      title: "University Partnerships",
      description:
        "Facilitate seamless collaboration between universities and partner companies, aligning academic goals with real-world experience.",
    },
    {
      icon: Users,
      title: "Startup Talent Acquisition",
      description:
        "Help startups attract and nurture top talent through well-structured internship programs, turning interns into valuable team members.",
    },
    {
      icon: Briefcase,
      title: "Remote Internship Management",
      description:
        "Enable smooth coordination of remote internships, ensuring clear communication and progress tracking across distributed teams.",
    },
  ];
</script>

<section class="bg-white py-16">
  <div class="container max-w-6xl mx-auto px-4">
    <div class="flex flex-col md:flex-row">
      <div class="md:w-5/12 md:pr-12 mb-8 md:mb-0">
        <h2 class="text-4xl font-bold mb-4">Usecase</h2>
        <p class="text-xl text-gray-600 mb-6">
          Transforming internship management across industries.
        </p>
        <a href="/" class="text-green-600 hover:underline font-semibold"
          >Explore all use cases &rarr;</a
        >
      </div>
      <div class="md:w-7/12">
        <div class="space-y-8">
          {#each useCases as useCase}
            <div class="flex items-start">
              <div class="flex-shrink-0 mr-4">
                <div
                  class="w-10 h-10 flex items-center justify-center rounded-full bg-gray-100"
                >
                  <svelte:component
                    this={useCase.icon}
                    class="w-5 h-5 text-gray-500"
                  />
                </div>
              </div>
              <div>
                <h3 class="text-xl font-semibold mb-2">{useCase.title}</h3>
                <p class="text-gray-600">{useCase.description}</p>
              </div>
            </div>
          {/each}
        </div>
      </div>
    </div>
  </div>
</section>

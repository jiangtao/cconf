import {
  Shield,
  GitBranch,
  Zap,
  Globe,
  Package,
  Settings,
} from "lucide-react"

const features = [
  {
    icon: Zap,
    title: "One-Command Backup",
    description:
      "Single command backs up all your Claude Code configurations to Git.",
  },
  {
    icon: GitBranch,
    title: "Version Control",
    description:
      "All configs stored in Git with full history. Roll back anytime.",
  },
  {
    icon: Shield,
    title: "Secure",
    description:
      "API tokens automatically excluded. Your secrets never leave your machine.",
  },
  {
    icon: Package,
    title: "Cross-Platform",
    description:
      "Single binary works on macOS, Linux, and Windows. No dependencies.",
  },
  {
    icon: Globe,
    title: "Internationalization",
    description: "Full support for English and 中文 (Chinese) interfaces.",
  },
  {
    icon: Settings,
    title: "Project Discovery",
    description:
      "Auto-scans your common directories to find all Claude Code projects.",
  },
]

export function Features() {
  return (
    <section className="border-t bg-slate-50 py-16 sm:py-24">
      <div className="container mx-auto px-4">
        <div className="mb-8 text-center sm:mb-12">
          <h2 className="mb-4 text-2xl font-bold text-slate-900 sm:text-3xl">
            Everything you need
          </h2>
          <p className="text-base text-slate-600 sm:text-lg">
            Built specifically for Claude Code power users
          </p>
        </div>
        <div className="grid gap-4 sm:gap-6 md:grid-cols-2 lg:grid-cols-3">
          {features.map((feature) => (
            <div
              key={feature.title}
              className="rounded-lg border bg-white p-4 shadow-sm sm:p-6"
            >
              <feature.icon className="mb-3 h-6 w-6 text-blue-600 sm:mb-4 sm:h-8 sm:w-8" />
              <h3 className="mb-2 text-base font-semibold text-slate-900 sm:text-lg">
                {feature.title}
              </h3>
              <p className="text-sm text-slate-600 sm:text-base">{feature.description}</p>
            </div>
          ))}
        </div>
      </div>
    </section>
  )
}

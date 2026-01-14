'use client'

import {
  Shield,
  GitBranch,
  Zap,
  Globe,
  Package,
  Settings,
} from "lucide-react"
import { useLanguage } from "@/contexts/language-context"
import { translations } from "@/lib/translations"

const featureIcons = {
  oneCommand: Zap,
  versionControl: GitBranch,
  secure: Shield,
  crossPlatform: Package,
  i18n: Globe,
  projectDiscovery: Settings,
}

export function Features() {
  const { language } = useLanguage()
  const t = translations[language]

  return (
    <section className="border-t bg-slate-50 py-16 sm:py-24">
      <div className="container mx-auto px-4">
        <div className="mb-8 text-center sm:mb-12">
          <h2 className="mb-4 text-2xl font-bold text-slate-900 sm:text-3xl">
            {t.features.title}
          </h2>
          <p className="text-base text-slate-600 sm:text-lg">
            {t.features.subtitle}
          </p>
        </div>
        <div className="grid gap-4 sm:gap-6 md:grid-cols-2 lg:grid-cols-3">
          {Object.entries(t.features.items).map(([key, value]) => {
            const Icon = featureIcons[key as keyof typeof featureIcons]
            return (
              <div
                key={key}
                className="rounded-lg border bg-white p-4 shadow-sm sm:p-6"
              >
                <Icon className="mb-3 h-6 w-6 text-blue-600 sm:mb-4 sm:h-8 sm:w-8" />
                <h3 className="mb-2 text-base font-semibold text-slate-900 sm:text-lg">
                  {value.title}
                </h3>
                <p className="text-sm text-slate-600 sm:text-base">{value.description}</p>
              </div>
            )
          })}
        </div>
      </div>
    </section>
  )
}

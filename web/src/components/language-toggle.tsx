'use client'

import { Languages } from 'lucide-react'
import { useLanguage } from '@/contexts/language-context'

export function LanguageToggle() {
  const { language, setLanguage } = useLanguage()

  return (
    <button
      type="button"
      onClick={() => setLanguage(language === 'en' ? 'zh' : 'en')}
      className="rounded-full p-2 text-slate-600 hover:bg-slate-100 transition-colors"
      aria-label="Toggle language"
      title={language === 'en' ? '切换到中文' : 'Switch to English'}
    >
      <Languages className="h-5 w-5" />
      <span className="ml-2 text-sm font-medium">{language === 'en' ? 'EN' : '中文'}</span>
    </button>
  )
}

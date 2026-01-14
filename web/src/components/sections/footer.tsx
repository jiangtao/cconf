import { Github } from "lucide-react"

export function Footer() {
  const currentYear = new Date().getFullYear()

  return (
    <footer className="border-t bg-slate-900 py-8 text-slate-300 sm:py-12">
      <div className="container mx-auto px-4">
        <div className="flex flex-col items-center justify-between gap-4 text-center md:flex-row md:text-left">
          <p className="text-xs sm:text-sm">© {currentYear} ccconfig. MIT License.</p>
          <div className="flex items-center gap-4 sm:gap-6">
            <a
              href="https://github.com/jiangtao/cc-config"
              className="flex items-center gap-2 hover:text-white"
              target="_blank"
              rel="noopener noreferrer"
            >
              <Github className="h-4 w-4 sm:h-5 sm:w-5" />
              <span className="text-sm">GitHub</span>
            </a>
            <a
              href="https://github.com/jiangtao/cc-config/blob/main/README.md"
              className="text-sm hover:text-white"
              target="_blank"
              rel="noopener noreferrer"
            >
              Documentation
            </a>
          </div>
        </div>
      </div>
    </footer>
  )
}

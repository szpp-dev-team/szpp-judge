/** https://github.com/mswjs/msw/discussions/1231 */
const ignoredPathnames = [
  "/_next/static/development",
  "/_next/static/webpack",
];

async function initMocks() {
  if (typeof window === "undefined") {
    // If you wanna use server side mock, refer to:
    // https://github.com/vercel/next.js/tree/f7baa56792cf8566954f762b961260a00c476994/examples/with-msw
  } else {
    const { worker } = await import("./browser");
    worker.start({
      onUnhandledRequest: (req, print) => {
        if (req.url.protocol === "chrome-extension:") {
          return;
        }
        if (ignoredPathnames.some(pathname => req.url.pathname.startsWith(pathname))) {
          return;
        }
        print.warning();
      },
    });
  }
}

initMocks();

export { };

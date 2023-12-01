const menuLinks = document.querySelectorAll("#aside-menu ul li a");
menuLinks.forEach((link) => {
  const pathname = location.pathname;
  const linkHref = link.getAttribute("href");
  if (pathname.includes(linkHref)) {
    link.classList.add("active");
  }
});
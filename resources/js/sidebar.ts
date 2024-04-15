const sidebar = document.getElementById('sidebar');

if (sidebar) {
  const toggleSidebarMobile = (sidebar: HTMLElement, sidebarBackdrop: HTMLElement, toggleSidebarMobileHamburger: HTMLElement, toggleSidebarMobileClose: HTMLElement) => {
    sidebar.classList.toggle('hidden');
    sidebarBackdrop.classList.toggle('hidden');
    toggleSidebarMobileHamburger.classList.toggle('hidden');
    toggleSidebarMobileClose.classList.toggle('hidden');
  }
  
  const toggleSidebarMobileEl = document.getElementById('toggleSidebarMobile');
  const sidebarBackdrop = document.getElementById('sidebarBackdrop');
  const toggleSidebarMobileHamburger = document.getElementById('toggleSidebarMobileHamburger');
  const toggleSidebarMobileClose = document.getElementById('toggleSidebarMobileClose');
  const toggleSidebarMobileSearch = document.getElementById('toggleSidebarMobileSearch');

  if (sidebarBackdrop && toggleSidebarMobileHamburger && toggleSidebarMobileClose) {
    toggleSidebarMobileSearch && toggleSidebarMobileSearch.addEventListener('click', () => {
      sidebar.classList.add('flex');
      toggleSidebarMobile(sidebar, sidebarBackdrop, toggleSidebarMobileHamburger, toggleSidebarMobileClose);
    });
    
    toggleSidebarMobileEl && toggleSidebarMobileEl.addEventListener('click', () => {
      sidebar.classList.add('flex');
      toggleSidebarMobile(sidebar, sidebarBackdrop, toggleSidebarMobileHamburger, toggleSidebarMobileClose);
    });
    
    sidebarBackdrop.addEventListener('click', () => {
      toggleSidebarMobile(sidebar, sidebarBackdrop, toggleSidebarMobileHamburger, toggleSidebarMobileClose);
    });
  }
  
}

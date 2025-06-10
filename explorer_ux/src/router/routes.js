const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', 
	component: () => import('pages/DashboardPage.vue'),
        meta: { title: 'Dashboard - MCP Explorer' }
      },
  
      { path: 'server_ops',
	component: () => import('pages/ServerBuilder.vue'),
	meta: { title: 'MCP Server Builder - MCP Explorer' }
      }
    ]
  },

  

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes

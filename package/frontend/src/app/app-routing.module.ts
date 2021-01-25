import { NgModule } from '@angular/core';
import { PreloadAllModules, RouterModule, Routes } from '@angular/router';

import { AuthGuard } from './core/guards/auth.guard';
import { NoAuthGuard } from './core/guards/no-auth.guard';

const routes: Routes = [
    {
        path: 'auth',
        loadChildren: () =>
            import('@pages/auth/auth.module').then((m) => m.AuthPageModule),
        canActivate: [NoAuthGuard],
    },
    {
        path: 'home',
        loadChildren: () =>
            import('./pages/home/home.module').then((m) => m.HomePageModule),
        canActivate: [AuthGuard],
    },
    {
        path: 'movies',
        loadChildren: () =>
            import('./pages/movies/movies.module').then(
                (m) => m.MoviesPageModule
            ),
        canActivate: [AuthGuard],
    },
    {
        path: '**',
        pathMatch: 'full',
        redirectTo: 'auth',
    },
];

@NgModule({
    imports: [
        RouterModule.forRoot(routes, { preloadingStrategy: PreloadAllModules }),
    ],
    exports: [RouterModule],
})
export class AppRoutingModule {}

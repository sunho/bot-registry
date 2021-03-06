import { BotResolverService as BotResolver } from './services/bot-resolve.service';
import { VolumeResolverService as VolumeResolver } from './services/volume-resolve.service';
import { FirstBotRedirectService as FirstBotRedirect } from './services/first-bot-resolve.service';
import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { DashComponent } from './pages/dash/dash.component';
import { HomeComponent } from './pages/home/home.component';
import { ConfigComponent } from './pages/config/config.component';
import { VolumeComponent } from './pages/volume/volume.component';
import { VolumeDetailComponent } from './pages/volume-detail/volume-detail.component';

const routes: Routes = [
  {
    path: '',
    component: DashComponent,
    canActivate: [FirstBotRedirect],
  },
  {
    path: ':id',
    component: DashComponent,
    resolve: {
      bot: BotResolver,
    },
    children: [
      { path: '', redirectTo: 'home', pathMatch: 'full' },
      { path: 'home', component: HomeComponent },
      { path: 'volume', component: VolumeComponent },
      { path: 'volume/:name', resolve: {
        volume: VolumeResolver
      }, component: VolumeDetailComponent },
      { path: 'config', component: ConfigComponent },
    ],
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule],
})
export class DashBoardRoutingModule {}

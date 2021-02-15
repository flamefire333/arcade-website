import {BrowserModule} from '@angular/platform-browser';
import {HttpClientModule} from '@angular/common/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations'
import {MatAutocompleteModule} from "@angular/material/autocomplete";
import {MatIconModule} from "@angular/material/icon";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatInputModule} from "@angular/material/input";
import {MatSelectModule} from "@angular/material/select";
import {MatOptionModule} from "@angular/material/core";
import {MatButtonModule} from "@angular/material/button";
import {MatTreeModule} from "@angular/material/tree";
import {ReactiveFormsModule} from "@angular/forms";
import {NgModule} from '@angular/core';
import {AppComponent} from './app.component';
import {LoginComponent} from './login/login.component';
import { LobbyComponent } from './lobby/lobby.component';
import { RouterModule, Routes } from '@angular/router';
import { Fe3hauComponent } from './fe3hau/fe3hau.component';
import { BeanPreviewComponent } from './bean-preview/bean-preview.component';
import { FeBeanPreviewComponent } from './fe-bean-preview/fe-bean-preview.component'
import {MatDividerModule} from "@angular/material/divider";
import { FeBeanListComponent } from './fe-bean-list/fe-bean-list.component';
import { VisualNovelComponent } from './visual-novel/visual-novel.component';
import { DiscordMappingsAddComponent } from './discord-mappings-add/discord-mappings-add.component';
import { MafiaComponent } from './mafia/mafia.component';
import { ChatComponent } from './chat/chat.component';
import {MatTooltipModule} from "@angular/material/tooltip";
import { GameComponent } from './game/game.component';
import { IconWithTooltipComponent } from './icon-with-tooltip/icon-with-tooltip.component';

const routes: Routes = [
  { path: 'discord-mapping', component: DiscordMappingsAddComponent },
  { path: 'login', component: LoginComponent },
  { path: 'game/:game/:username', component: GameComponent },
  { path: 'display', component: FeBeanListComponent },
  { path: 'visual-novel', component: VisualNovelComponent },
  { path: '', redirectTo: '/login', pathMatch: 'full' },
];

@NgModule
({
  declarations: [
    AppComponent,
    LoginComponent,
    LobbyComponent,
    Fe3hauComponent,
    BeanPreviewComponent,
    FeBeanPreviewComponent,
    FeBeanListComponent,
    VisualNovelComponent,
    DiscordMappingsAddComponent,
    MafiaComponent,
    ChatComponent,
    GameComponent,
    IconWithTooltipComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatIconModule,
    ReactiveFormsModule,
    MatAutocompleteModule,
    MatInputModule,
    MatFormFieldModule,
    MatOptionModule,
    MatSelectModule,
    MatTreeModule,
    MatButtonModule,
    RouterModule.forRoot(routes),
    MatDividerModule,
    MatTooltipModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {

}

import { Component, OnInit } from '@angular/core';
import {FormBuilder, FormControl, FormGroup, Validators} from "@angular/forms";
import {HttpClient} from "@angular/common/http";
import {Router} from "@angular/router";
import {MatAutocomplete} from "@angular/material/autocomplete";
import {MatTree, MatTreeNestedDataSource} from "@angular/material/tree";
import {map, startWith} from "rxjs/operators";
import {Observable} from "rxjs";
import {NestedTreeControl} from "@angular/cdk/tree";

interface PersonNode {
  name: string;
  children?: PersonNode[];
}

interface PersonNodeTemp {
  name: string;
  children?: number[];
  parent: string;
}

const TREE_DATA: PersonNode[] = [];

@Component({
  selector: 'app-discord-mappings-add',
  templateUrl: './discord-mappings-add.component.html',
  styleUrls: ['./discord-mappings-add.component.css']
})
export class DiscordMappingsAddComponent implements OnInit {

  loginForm : FormGroup;
  selected: string;
  parents: string[] = [];
  filteredParents: Observable<string[]>;
  statusText = "";
  treeControl = new NestedTreeControl<PersonNode>(node => node.children);
  dataSource = new MatTreeNestedDataSource<PersonNode>();

  constructor(private http: HttpClient, private formBuilder: FormBuilder, private router: Router) {
    this.loginForm = this.formBuilder.group({
      name: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(20)]],
      parent: ['', [Validators.required, Validators.minLength(2), Validators.maxLength(20)]]
    });
    this.dataSource.data = TREE_DATA;
    this.refreshParents()
    this.filteredParents = this.loginForm.controls["parent"].valueChanges.pipe(
      startWith(''),
      map(value => this.filterNames(value))
    )
  }

  ngOnInit(): void {
    setInterval(() => {
      this.refreshParents();
    }, 5000)
    const url : string = "http://jaketheduck.com/discord-mapping-api/api.php?type=alldata";
    this.http.get(url).subscribe((data : {}) => {
      console.log("Data: " + data);
      console.log("Status: " + data["status"]);
      if(data["status"] == 0) {
        const basis : PersonNodeTemp[] = data["players"].map(p => {
          return {name: p["name"], parent: p["parent"], children: []};
        });
        basis.push({name: "?", parent: "?", children: []});
        const qindex = basis.length - 1;
        basis.forEach(b => {
          let parentIndex = basis.findIndex(p => p.name == b.parent);
          const myIndex = basis.findIndex(p => p.name == b.name);
          if(parentIndex == -1) {
            basis.push({name: b.parent, parent: "?", children: []})
            parentIndex = basis.length - 1;
            basis[qindex].children.push(parentIndex);
          }
          if(myIndex != parentIndex) {
            basis[parentIndex].children.push(myIndex);
          }
        })
        console.log(basis);
        this.dataSource.data = basis.filter(b => b.name == b.parent).map(b => {return this.qqq(basis.findIndex(p => p.name == b.name), basis);});
        console.log(this.dataSource.data);
      }
    })
  }

  refreshParents() : void {
    const url : string = "http://jaketheduck.com/discord-mapping-api/api.php?type=data";
    this.http.get(url).subscribe((data : {}) => {
      console.log("Data: " + data);
      console.log("Status: " + data["status"]);
      if(data["status"] == 0) {
        this.parents = data["players"];
      }
    })
  }

  filterNames(name: string) : string[] {
    console.log("I CRY");
   return this.parents.filter(option => option.toLowerCase().includes(name.toLowerCase()))
  }

  submit(name: string, parent: string) {
    const url : string = "http://jaketheduck.com/discord-mapping-api/api.php?type=update&name=" + name.toLowerCase() + "&parent=" + parent.toLowerCase();
    this.http.get(url).subscribe((data : {}) => {
      console.log("Data: " + data);
      console.log("Status: " + data["status"]);
      if(data["status"] == 0) {
        this.statusText = name + " was added successfully";
      } else {
        this.statusText = "request was not successful with code: " + data["status"];
      }
    })
  }

  qqq(i : number, basis: PersonNodeTemp[]) : PersonNode {
    return {name: basis[i].name, children: basis[i].children.map(j => this.qqq(j, basis))}
  }

  hasChild = (_: number, node: PersonNode) => !!node.children && node.children.length > 0;
}

